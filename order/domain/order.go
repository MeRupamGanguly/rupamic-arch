package domain

import "time"

type OrderType string

const (
	Long  OrderType = "Long"
	Short OrderType = "Short"
)

type OrderPriceType string

const (
	Market OrderPriceType = "Market"
	Limit  OrderPriceType = "Limit"
)

type Order struct {
	Id             string
	UserId         string
	Symbol         string
	Qty            int
	OrderType      OrderType
	OrderPriceType OrderPriceType
	LimitPrice     float64
	BuyPrice       float64
	BuyTimestamp   time.Time
	SellPrice      float64
	SellTimestamp  time.Time
	PNL            float64
	Completed      bool
	Notified       bool
}

type Orders struct {
	UserId    string
	Orders    []Order
	Aware     bool
	OveralPNL float64
}
