package infrastructure

import (
	"context"
	"fmt"
	"time"

	"EventAgent.Consumer/configuration"
	"github.com/segmentio/kafka-go"
)

var partition = 0

type KafkaBase struct {
	Configuration configuration.Configuration
}

func (k KafkaBase) Consumer(ctx context.Context, topic string) error {
	fmt.Println("Consumer init for: " + topic)

	if k.Configuration.BrokerAddress == "" {
		panic("configuration.BrokderAddress can not be nil")
	}
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{k.Configuration.BrokerAddress},
		GroupID: "consumer-group-id",
		Topic:   topic,
		Dialer:  dialer,
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			return err
		}
		fmt.Println(string(msg.Value))
	}
}
