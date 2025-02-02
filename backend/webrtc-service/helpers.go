package main

import (
	"log"
	"os"
	"reflect"

	"github.com/IBM/sarama"
	"github.com/Slimo300/Microservices-Videocall-and-Chat-App/backend/lib/events"
	"github.com/Slimo300/Microservices-Videocall-and-Chat-App/backend/lib/msgqueue"
	"github.com/Slimo300/Microservices-Videocall-and-Chat-App/backend/lib/msgqueue/kafka"
)

// kafkaSetup starts kafka EventEmiter and EventListener
func kafkaSetup(brokerAddreses []string) (emiter msgqueue.EventEmiter, listener msgqueue.EventListener, err error) {
	brokerConf := sarama.NewConfig()
	brokerConf.ClientID = "webrtcService"
	brokerConf.Version = sarama.V2_3_0_0
	brokerConf.Producer.Return.Successes = true
	client, err := sarama.NewClient(brokerAddreses, brokerConf)
	if err != nil {
		return nil, nil, err
	}

	// initializing emiter
	emiter, err = kafka.NewKafkaEventEmiter(client, log.New(os.Stdout, "[ emiter ]: ", log.Flags()))
	if err != nil {
		return nil, nil, err
	}

	// initializing dbListener
	dbListenerMapper := msgqueue.NewDynamicEventMapper()
	if err := dbListenerMapper.RegisterTypes(
		reflect.TypeOf(events.GroupDeletedEvent{}),
		reflect.TypeOf(events.MemberCreatedEvent{}),
		reflect.TypeOf(events.MemberUpdatedEvent{}),
		reflect.TypeOf(events.MemberDeletedEvent{}),
	); err != nil {
		return nil, nil, err
	}

	listener, err = kafka.NewConsumerGroupEventListener(client, "webrtc-service", dbListenerMapper, &kafka.ListenerOptions{
		Logger: log.New(os.Stdout, "[ listener ]: ", log.Flags()),
	})
	if err != nil {
		return nil, nil, err
	}

	return emiter, listener, nil
}
