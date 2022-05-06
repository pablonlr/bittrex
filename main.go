package main

import (
	"fmt"

	"github.com/pablonlr/bittrex/client"
)

func main() {
	bclient := client.NewClient("", "")
	resp, err := bclient.OrderBook("ALGO-USDT")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

}
