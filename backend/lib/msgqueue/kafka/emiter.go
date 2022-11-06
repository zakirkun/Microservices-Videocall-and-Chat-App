package kafka

import (
	"encoding/json"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/Slimo300/MicroservicesChatApp/backend/lib/msgqueue"
)

type kafkaEventEmiter struct {
	producer sarama.SyncProducer
}

type kafkaMessage struct {
	EventName string      `json:"eventName"`
	Payload   interface{} `json:"payload"`
}

func NewKafkaEventEmiter(client sarama.Client) (msgqueue.EventEmitter, error) {
	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		return nil, err
	}

	return &kafkaEventEmiter{
		producer: producer,
	}, nil
}

func (k *kafkaEventEmiter) Emit(event msgqueue.Event) error {
	jsonBody, err := json.Marshal(kafkaMessage{
		EventName: event.EventName(),
		Payload:   event,
	})
	if err != nil {
		return err
	}

	topic := strings.Split(event.EventName(), ".")[0]

	_, _, err = k.producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(jsonBody),
	})
	return err
}
