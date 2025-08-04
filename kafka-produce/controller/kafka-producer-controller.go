package controller

import (
	core_model "confluence-payment/core-internal/model"
	"confluence-payment/kafka-produce/model"
	"confluence-payment/kafka-produce/service"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type KafkaProduceController struct {
	Validate            *validator.Validate
	KafkaProduceService *service.KafkaProduceService
}

func NewKafkaProducerController(KafkaProduceService *service.KafkaProduceService, validate *validator.Validate) *KafkaProduceController {

	return &KafkaProduceController{
		Validate:            validate,
		KafkaProduceService: KafkaProduceService,
	}
}

// This controller is only for testing or maintaining data
func (k *KafkaProduceController) Produce(g *gin.Context) {
	res := core_model.CoreResponseObject{}
	body := model.KafkaMessage{}
	topic := g.Param("topic")

	err := g.ShouldBindJSON(&body)
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		g.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err = k.Validate.Struct(body)
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		g.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	statusCode, err := k.KafkaProduceService.Produce(context.Background(), topic, body)
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		g.AbortWithStatusJSON(statusCode, res)
		return
	}

	res.Succeeded = true
	g.JSON(statusCode, res)
}
