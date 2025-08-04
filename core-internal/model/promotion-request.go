package model

import (
	"time"
)

type CorePromotion struct {
	CoreModel
	Name              string        `json:"name" validate:"required"`
	PromotionType     string        `json:"promotion_type"`
	Code              string        `json:"code"`
	StartTime         time.Time     `json:"start_time" validate:"required"`
	EndTime           time.Time     `json:"end_time" validate:"required"`
	ShopID            string        `json:"shop_id"`
	UsageQuantity     int           `json:"usage_quantity"`
	UsageLimitPerUser int           `json:"usage_limit_per_user"`
	Products          []CoreProduct `json:"products" validate:"dive"`
}

type CoreProduct struct {
	CoreModel
	PromotionID     string               `json:"promotion_id"`
	SKU             string               `json:"sku" validate:"required"`
	Name            string               `json:"name" validate:"required"`
	PurchaseLimit   int                  `json:"purchase_limit"`
	ProductVariants []CoreProductVariant `json:"product_variants" validate:"dive"`
}

type CoreProductVariant struct {
	CoreModel
	PromotionID          string  `json:"promotion_id"`
	ProductID            string  `json:"product_id"`
	SKU                  string  `json:"sku" validate:"required"`
	Name                 string  `json:"name" validate:"required"`
	DiscountedPrice      float64 `json:"discounted_price" validate:"required,gte=0"`
	DiscountedPercentage float64 `json:"discounted_percentage"`
	StockLimit           int     `json:"stock_limit"`
	IsActive             bool    `json:"is_active" validate:"required"`
}
