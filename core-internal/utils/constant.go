package utils

import "time"

const (
	XPlatformKey = "x-platform-key"
	XShopId      = "x-shop-id"
)

const (
	DefaultPage  = "1"
	DefaultSize  = "10"
	DefaultOrder = "id DESC"
)

const (
	TopicCreateBulkDiscount = "topic-create-bulk-discount"
)

const (
	PaymentTableName = "\"payment\".payments"
)

const (
	PromotionTypeDiscount        = "discount"
	PromotionTypeFlashSale       = "flashsale"
	PromotionTypeAddonDeal       = "addon-deal"
	PromotionTypeBundleDeal      = "bundle-deal"
	PromotionTypeVoucher         = "voucher"
	PromotionTypeVoucherShipping = "voucher-shipping"
)

var (
	DefaultContextTimeOut = time.Second * 60000
	DefaultRedisTimeOut   = time.Second * 10000
)

// Insert
var PaymentColumnsListForInsert = []string{"uuid", "order_id", "shop_id", "customer_id", "amount"}
var PaymentColumnsListForSelect = []string{"uuid", "created_at", "updated_at", "deleted_at", "order_id", "shop_id", "customer_id", "amount"}