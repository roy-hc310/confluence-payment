package kafka_producer

import (
	"confluence-payment/kafka-produce/controller"

	"github.com/gin-gonic/gin"
)

func KafkaProduceRoute(KafkaProduceRoute *gin.RouterGroup, controller *controller.KafkaProduceController) {

	KafkaProduceRoute.POST("/:topic", controller.Produce)
}
