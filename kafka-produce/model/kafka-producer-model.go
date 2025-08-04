package model

import "time"

type KafkaMessage struct {
	Retry           *int        `json:"retry"`
	Data            interface{} `json:"data"`
	Error           *string     `json:"error"`
	LastProcessTime *time.Time  `json:"last_process_time"`
}