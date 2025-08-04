package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Middleware(g *gin.Context) {
	shopId := g.Request.Header.Get(XShopId)
	if shopId == "" {
		g.AbortWithStatusJSON(http.StatusUnauthorized, "Missing x-shop-id")
		return
	}
	g.Set(XShopId, shopId)
	g.Next()
}