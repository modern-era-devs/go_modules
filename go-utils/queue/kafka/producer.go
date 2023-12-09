package kafka

import (
	confluentKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
)

func SetupProducerConnection(cfg ProducerConfig) (*confluentKafka.Producer, error) {
	return confluentKafka.NewProducer(&confluentKafka.ConfigMap{
		"bootstrap.servers": cfg.BootstrapServers,
		"client.id":         cfg.ClientID,
		"acks":              cfg.Ack,
	})
}

func Produce(prodConn *confluentKafka.Producer, topic string, value []byte) error {
	deliveryChan := make(chan confluentKafka.Event, 1)
	defer close(deliveryChan)

	// TODO check the feasibility of producing messages in goroutines
	err := prodConn.Produce(&confluentKafka.Message{
		TopicPartition: confluentKafka.TopicPartition{Topic: &topic, Partition: confluentKafka.PartitionAny},
		Value:          value,
	}, deliveryChan)

	if err != nil {
		return errors.Wrap(err, "error while producing message")
	}

	chanOut := <-deliveryChan
	deliveryReport := chanOut.(*confluentKafka.Message)
	return deliveryReport.TopicPartition.Error
}
