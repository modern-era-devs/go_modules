package kafka

import (
	"context"
	"fmt"
	confluentKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
	"os"
)

func SetupConsumerConnection(cfg ConsumerConfig) (*confluentKafka.Consumer, error) {
	consumer, err := confluentKafka.NewConsumer(&confluentKafka.ConfigMap{
		"bootstrap.servers": cfg.BootstrapServers,
		"group.id":          cfg.ConsumerGroupID,
		"auto.offset.reset": cfg.AutoOffsetReset,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Error while creating kafka consumer")
	}
	err = consumer.Subscribe(cfg.Topic, nil)

	if err != nil {
		return nil, errors.Wrap(err, "Error while subscribing to topic: "+cfg.Topic)
	}

	return consumer, err
}

func StartConsumption(_ context.Context, consumer *confluentKafka.Consumer, pollTimeout int, exec func([]byte) error) error {
	run := true
	var err error
	for run == true {
		ev := consumer.Poll(pollTimeout)
		switch e := ev.(type) {
		case *confluentKafka.Message:
			// app specific processing here
			go func() {
				err := exec(e.Value)
				if err != nil {
					fmt.Fprintf(os.Stdout, "%% Error while executing message: %v\n", e)
				}
			}()
		case confluentKafka.Error:
			fmt.Fprintf(os.Stdout, "%% Error while reading message: %v\n", e)
			err = e
			run = false
		}
	}
	//closing consumer connection
	consumer.Close()
	return err
}
