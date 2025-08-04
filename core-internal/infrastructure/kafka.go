package infrastructure

import (
	"confluence-payment/core-internal/utils"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/twmb/franz-go/pkg/kgo"
)

type KafkaInfra struct {
	Client *kgo.Client
}

func NewKafkaInfra() *KafkaInfra {

	client, err := kgo.NewClient(
		kgo.AllowAutoTopicCreation(),
		kgo.SeedBrokers(strings.Split(utils.GlobalEnv.KafkaHost, ",")...),
		kgo.ConsumerGroup(utils.GlobalEnv.KafkaConsumerGroup),
		kgo.ConsumeTopics(
			utils.TopicCreateBulkDiscount,
		),
	)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	return &KafkaInfra{
		Client: client,
	}
}
