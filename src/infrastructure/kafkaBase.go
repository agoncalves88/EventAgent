package infrastructure

import (
	"context"
	"log"

	"EventAgent.Consumer/configuration"
	"github.com/segmentio/kafka-go"
)

var partition = 0

type KafkaBase struct {
	Configuration configuration.Configuration
}

func (k KafkaBase) GetKafkaConnection(topic string, context context.Context) *kafka.Conn {
	conn, err := kafka.DialLeader(context, "tcp", k.Configuration.BrokerAddress, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	return conn
}
