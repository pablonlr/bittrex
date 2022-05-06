package types

import (
	"strconv"

	"github.com/pablonlr/exchange"
)

func ConvertBook(bookstring *OrderBookString) (*exchange.OrderBook, error) {
	book := &exchange.OrderBook{}
	bid, err := orderStoF(bookstring.Bid)
	if err != nil {
		return nil, err
	}
	ask, err := orderStoF(bookstring.Ask)
	if err != nil {
		return nil, err
	}
	book.Ask = ask
	book.Bid = bid
	return book, nil

}

func orderStoF(stringOrders []OrderString) ([]exchange.Order, error) {
	orderfloat := []exchange.Order{}
	for _, v := range stringOrders {
		quanty, err := strconv.ParseFloat(v.Quantity, 64)
		if err != nil {
			return nil, err
		}
		rate, err := strconv.ParseFloat(v.Rate, 64)
		if err != nil {
			return nil, err
		}
		orderfloat = append(orderfloat, exchange.Order{Quantity: quanty, Price: rate})
	}
	return orderfloat, nil
}
