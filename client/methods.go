package client

import (
	"encoding/json"

	"github.com/pablonlr/bittrex/types"
	"github.com/pablonlr/exchange"
)

func (client *Client) OrderBook(marketSymbol string) (*exchange.OrderBook, error) {
	bookString, err := client.orderBookString(marketSymbol)
	if err != nil {
		return nil, err
	}
	book, err := types.ConvertBook(bookString)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (client *Client) orderBookString(marketSymbol string) (*types.OrderBookString, error) {
	resp, err := client.getReq("markets", marketSymbol, "orderbook")
	if err != nil {
		return nil, err
	}
	bookString := &types.OrderBookString{}
	err = json.Unmarshal(resp, bookString)
	if err != nil {
		return nil, err
	}
	return bookString, nil
}

func (client *Client) Balances() ([]types.BalanceResponse, error) {
	resp, err := client.Request(nil, "GET", "balances")
	if err != nil {
		return nil, err
	}
	balances := []types.BalanceResponse{}
	err = json.Unmarshal(resp, &balances)
	if err != nil {
		return nil, err
	}
	return balances, nil

}

func (client *Client) BalanceOfCoin(coinSymbol string) (*types.BalanceResponse, error) {
	resp, err := client.Request(nil, "GET", "balances", coinSymbol)
	if err != nil {
		return nil, err
	}
	balance := &types.BalanceResponse{}
	err = json.Unmarshal(resp, balance)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

// Orders

func (client *Client) OpenOrders() ([]types.OpenOrder, error) {
	resp, err := client.Request(nil, "GET", "orders", "open")
	if err != nil {
		return nil, err
	}
	orders := []types.OpenOrder{}
	err = json.Unmarshal(resp, &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (client *Client) GetOpenOrder(orderID string) (*types.OpenOrder, error) {
	resp, err := client.Request(nil, "GET", "orders", orderID)
	if err != nil {
		return nil, err
	}
	order := &types.OpenOrder{}
	err = json.Unmarshal(resp, order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (client *Client) CancelAllOpenOrders() ([]types.CancelOrderResponse, error) {
	resp, err := client.Request(nil, "DELETE", "orders", "open")
	if err != nil {
		return nil, err
	}
	orders := []types.CancelOrderResponse{}
	err = json.Unmarshal(resp, &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (client *Client) CancelOpenOrder(orderID string) (*types.CancelOrderResponse, error) {
	resp, err := client.Request(nil, "DELETE", "orders", orderID)
	if err != nil {
		return nil, err
	}
	order := &types.CancelOrderResponse{}
	err = json.Unmarshal(resp, order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

/*
direction should be 'BUY' or 'SELL'
*/
func (client *Client) NewLimitOrder(marketSymbol string, direction string, amount float64, limitPrice float64) (*types.OpenOrder, error) {
	norder := &types.NOrder{
		MarketSymbol: marketSymbol,
		Direction:    direction,
		Type:         "LIMIT",
		Quantity:     amount,
		Limit:        limitPrice,
		TimeInForce:  "GOOD_TIL_CANCELLED",
	}
	jso, err := json.Marshal(norder)
	if err != nil {
		return nil, err
	}
	resp, err := client.Request(jso, "POST", "orders")
	if err != nil {
		return nil, err
	}
	order := &types.OpenOrder{}
	err = json.Unmarshal(resp, order)
	if err != nil {
		return nil, err
	}
	return order, nil

}
