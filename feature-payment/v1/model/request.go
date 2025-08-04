package model

import "confluence-payment/core-internal/model"

type Payment struct {
	model.CoreModel
	OrderID    string `json:"order_id"`
	ShopID     string `json:"shop_id"`
	CustomerID string `json:"customer_id"`
	Amount     float64  `json:"amount"`
}
