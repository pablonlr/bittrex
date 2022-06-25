package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const baseURL string = "https://api.bittrex.com/v3/"

type Client struct {
	http.Client
	API_KEY    string
	SECRET_KEY string
}

func NewClient(api_key, secret_key string, timeout time.Duration) *Client {
	return &Client{
		Client:     http.Client{Timeout: timeout},
		API_KEY:    api_key,
		SECRET_KEY: secret_key,
	}
}

func (client *Client) getReq(endpoints ...string) ([]byte, error) {
	url := baseURL + strings.Join(endpoints, "/")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return client.doReq(req)
}
func (client *Client) doReq(req *http.Request) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 {
		return nil, errors.New(string(body))
	}
	return body, nil
}

func signBody(body []byte, secretAPI string) string {
	hm := hmac.New(sha512.New, []byte(secretAPI))
	hm.Write(body)
	sha := hex.EncodeToString(hm.Sum(nil))
	return sha
}

func sha5(body []byte) string {
	payload := sha512.Sum512(body)
	return hex.EncodeToString(payload[:])

}

func timeStamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func (client *Client) RequestAndEncodeBody(body map[string]interface{}, requestMethod string, endpoints ...string) ([]byte, error) {
	contBody := []byte{}
	if body != nil {
		var err error
		contBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}
	return client.Request(contBody, requestMethod, endpoints...)

}

func (client *Client) Request(encodedBody []byte, requestMethod string, endpoints ...string) ([]byte, error) {
	url := baseURL + strings.Join(endpoints, "/")
	req, err := http.NewRequest(requestMethod, url, bytes.NewReader(encodedBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Api-Key", client.API_KEY)
	tim := strconv.FormatInt(timeStamp(), 10)
	req.Header.Set("Api-Timestamp", tim)
	siBod := sha5(encodedBody)
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Accept", "application/json")
	req.Header.Set("Api-Content-Hash", siBod)
	preSigApi := strings.Join([]string{tim, url, requestMethod, siBod, ""}, "")
	sigApi := signBody([]byte(preSigApi), client.SECRET_KEY)
	req.Header.Set("Api-Signature", sigApi)
	resp, err := client.doReq(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
