package core

import (

	health_v1 "confluence-payment/feature-health/v1"
	kafka_producer "confluence-payment/kafka-produce"
	payment_v1 "confluence-payment/feature-payment/v1"

	"github.com/gin-gonic/gin"
)

func Router(g *gin.Engine, application *Application) error {

	g.GET("", application.HealthV1Controller.Health)

	route := g.Group("/api")
	healthRoute := route.Group("/health")
	health_v1.HealthV1Route(healthRoute, application.HealthV1Controller)

	KafkaProduceRoute := route.Group("/kafka")
	kafka_producer.KafkaProduceRoute(KafkaProduceRoute, application.KafkaProducerController)

	paymentRoute := route.Group("/payment")
	payment_v1.PaymentV1Route(paymentRoute, application.PaymentV1Controller)

	return nil
}
