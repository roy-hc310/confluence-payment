package service

import (
	"confluence-payment/kafka-produce/model"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
)

type KafkaProduceService struct {
	Client *kgo.Client
}

func NewKafkaProduceService(client *kgo.Client) *KafkaProduceService {
	return &KafkaProduceService{
		Client: client,
	}
}

func (k *KafkaProduceService) Produce(ctx context.Context, topic string, data model.KafkaMessage) (statusCode int, err error) {
	msg, err := json.Marshal(data)
	if err != nil {
		return http.StatusBadRequest, err
	}

	record := &kgo.Record{
		Topic:     topic,
		Value:     msg,
		Timestamp: time.Now(),
	}

	err = k.Client.ProduceSync(ctx, record).FirstErr()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
