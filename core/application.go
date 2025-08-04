package core

import (
	"confluence-payment/core-internal/conf"
	"confluence-payment/core-internal/infrastructure"
	health_v1_controller "confluence-payment/feature-health/v1/controller"
	health_v1_service "confluence-payment/feature-health/v1/service"

	kafka_producer_controller "confluence-payment/kafka-produce/controller"
	kafka_producer_service "confluence-payment/kafka-produce/service"

	payment_v1_controller "confluence-payment/feature-payment/v1/controller"
	payment_v1_service          "confluence-payment/feature-payment/v1/service"

	"github.com/go-playground/validator/v10"
)

type Application struct {
	HealthV1Service    *health_v1_service.HealthService
	HealthV1Controller *health_v1_controller.HealthController

	KafkaProduceService     *kafka_producer_service.KafkaProduceService
	KafkaProducerController *kafka_producer_controller.KafkaProduceController

	PaymentV1Service    *payment_v1_service.PaymentService
	PaymentV1Controller *payment_v1_controller.PaymentController
	PaymentProtoService *payment_v1_service.PaymentProtoService

}

func NewApplication(env conf.Env) *Application {
	application := &Application{}

	postgresInfra := infrastructure.NewPostgresInfra()
	kafkaInfra := infrastructure.NewKafkaInfra()
	redisInfra := infrastructure.NewRedisInfra()
	validate := validator.New(validator.WithRequiredStructEnabled())
	otelInfra := infrastructure.NewOtelInfra()

	application.HealthV1Service = health_v1_service.NewHealthService()
	application.HealthV1Controller = health_v1_controller.NewHealthController(application.HealthV1Service)

	application.KafkaProduceService = kafka_producer_service.NewKafkaProduceService(kafkaInfra.Client)
	application.KafkaProducerController = kafka_producer_controller.NewKafkaProducerController(application.KafkaProduceService, validate)
	

	application.PaymentV1Service = payment_v1_service.NewPaymentService(postgresInfra, redisInfra, otelInfra)
	application.PaymentV1Controller = payment_v1_controller.NewPaymentController(application.PaymentV1Service)
	application.PaymentProtoService = payment_v1_service.NewPaymentProtoService(application.PaymentV1Service, postgresInfra)

	return application
}
