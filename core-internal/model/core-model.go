package model

import "github.com/jackc/pgx/v5/pgtype"

type CoreModel struct {
	ID        int              `json:"id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	DeletedAt pgtype.Timestamp `json:"deleted_at"`
	UUID      string           `json:"uuid"`
	TraceID   string           `json:"trace_id"`
}

type CoreQuery struct {
	Page   string `form:"page" json:"page"`
	Size   string `form:"size" json:"size"`
	Cursor string `form:"cursor" json:"cursor"`

	Sort   string `form:"sort" json:"sort"`
	Search string `form:"search" json:"search"`
	ShopID string `form:"shop_id" json:"shop_id"`
}

type Pagination struct {
	Page       string    `json:"page"`
	Size       string    `json:"size"`
	TotalItems string    `json:"total_items"`
	NextCursor string `json:"next_cursor"`
}

type CoreResponseObject struct {
	Data      interface{} `json:"data"`
	TraceID   string      `json:"trace_id"`
	Succeeded bool        `json:"succeeded"`
	Errors    []string    `json:"errors"`
}

type CoreResponseArray struct {
	CoreResponseObject
	Pagination
}
