package redpanda

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type Consumer interface {
	Start(handler func(key string, value []byte) error) error
	Close() error
}

type RedpandaConsumer struct {
	reader *kafka.Reader
}

func NewConsumer(brokers []string, groupID, topic string) *RedpandaConsumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       1,
		MaxBytes:       10e6,
		CommitInterval: 1 * time.Second,
	})

	return &RedpandaConsumer{reader: r}
}

func (c *RedpandaConsumer) Start(handler func(key string, value []byte) error) error {
	ctx := context.Background()

	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			log.Println("[RedpandaConsumer] Read error:", err)
			time.Sleep(time.Second)
			continue
		}

		if err := handler(string(msg.Key), msg.Value); err != nil {
			log.Println("[RedpandaConsumer] handler error:", err)
		}
	}
}

func (c *RedpandaConsumer) Close() error {
	return c.reader.Close()
}
