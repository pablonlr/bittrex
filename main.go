package main

import (
	"fmt"
	"time"

	"github.com/pablonlr/bittrex/client"
)

func main() {
	bclient := client.NewClient("", "", 5*time.Second)
	resp, err := bclient.OrderBook("ALGO-USDT")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

}
