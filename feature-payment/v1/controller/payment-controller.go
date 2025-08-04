package controller

import (
	core_model "confluence-payment/core-internal/model"
	"confluence-payment/feature-payment/v1/model"
	"confluence-payment/feature-payment/v1/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	PaymentService *service.PaymentService
}

func NewPaymentController(paymentService *service.PaymentService) *PaymentController {
	return &PaymentController{PaymentService: paymentService}
}

func (a *PaymentController) CreatePayment(g *gin.Context) {
	res := core_model.CoreResponseObject{}

	body := model.Payment{}
	err := g.ShouldBindJSON(&body)
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		g.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	data, traceID, statusCode, err := a.PaymentService.CreatePayment(g.Request.Context(), body)
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		g.AbortWithStatusJSON(statusCode, res)
		return
	}

	res.Data = data
	res.TraceID = traceID
	res.Succeeded = true
	g.JSON(statusCode, res)
}

func (a *PaymentController) DetailPayment(g *gin.Context) {
	res := core_model.CoreResponseObject{}

	id := g.Param("id")

	data, traceID, statusCode, err := a.PaymentService.DetailPayment(g.Request.Context(), id, model.Payment{})
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		g.AbortWithStatusJSON(statusCode, res)
		return
	}

	res.Data = data
	res.TraceID = traceID
	res.Succeeded = true
	g.JSON(statusCode, res)
}

func (a *PaymentController) ListPayment(g *gin.Context) {
	res := core_model.CoreResponseArray{}

	query := core_model.CoreQuery{}
	err := g.ShouldBindQuery(&query)
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		g.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	data, pagination, traceID, statusCode, err := a.PaymentService.ListPayment(g.Request.Context(), query)
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		g.AbortWithStatusJSON(statusCode, res)
		return
	}

	res.Data = data
	res.Pagination = pagination
	res.TraceID = traceID
	res.Succeeded = true
	g.JSON(statusCode, res)
}

