package v1

import (
	"confluence-payment/feature-payment/v1/controller"

	"github.com/gin-gonic/gin"
)

func PaymentV1Route(paymentRoute *gin.RouterGroup, controller *controller.PaymentController) {
	v1Route := paymentRoute.Group("/v1")

	v1Route.POST("", controller.CreatePayment)
	v1Route.GET("", controller.ListPayment)
	v1Route.GET("/:id", controller.DetailPayment)
}
