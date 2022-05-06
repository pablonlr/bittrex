package types

import (
	"github.com/pablonlr/exchange"
)

type OrderBookString struct {
	Bid []OrderString
	Ask []OrderString
}

type OrderString struct {
	Quantity string
	Rate     string
}

type Order exchange.Order
type OrderBook exchange.OrderBook

type BalanceResponse struct {
	CurrencySymbol string
	Total          string
	Available      string
	UpdatedAt      string
}

type CancelOrderResponse struct {
	ID         string
	StatusCode string
	Result     OpenOrder
}

type OpenOrder struct {
	ID            string
	MarketSymbol  string
	Direction     string
	Type          string
	Quantity      string
	Limit         string
	Ceiling       string
	TimeInForce   string
	ClientOrderId string
	FillQuantity  string
	Commission    string
	Proceeds      string
	Status        string
	CreatedAt     string
	UpdatedAt     string
	ClosedAt      string
	OrderToCancel OrderToCancelResponse
}

type OrderToCancelResponse struct {
	Type string
	ID   string
}

type NOrder struct {
	MarketSymbol  string  `json:"marketSymbol"`
	Direction     string  `json:"direction"`
	Type          string  `json:"type"`
	Quantity      float64 `json:"quantity"`
	Ceiling       float64 `json:"ceiling,omitempty"`
	Limit         float64 `json:"limit"`
	TimeInForce   string  `json:"timeInForce,omitempty"`
	ClientOrderId string  `json:"clientOrderId,omitempty"`
	UseAwards     bool    `json:"useAwards,omitempty"`
}
