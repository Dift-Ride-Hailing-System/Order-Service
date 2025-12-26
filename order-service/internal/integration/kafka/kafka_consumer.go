package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer interface {
	Consume(handler func(key string, value []byte) error) error
	Close() error
}

type KafkaReader struct {
	reader *kafka.Reader
}

func NewKafkaReader(brokers []string, groupID, topic string) *KafkaReader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          topic,
		CommitInterval: time.Second, // commit ทุก 1 วิ
		MinBytes:       1,
		MaxBytes:       10e6, // 10MB
	})
	return &KafkaReader{reader: r}
}

func (k *KafkaReader) Consume(handler func(key string, value []byte) error) error {
	ctx := context.Background()

	for {
		msg, err := k.reader.ReadMessage(ctx)
		if err != nil {
			log.Println("[KafkaConsumer] Read error:", err)
			time.Sleep(1 * time.Second)
			continue
		}

		if handlerErr := handler(string(msg.Key), msg.Value); handlerErr != nil {
			log.Println("[KafkaConsumer] Handler error:", handlerErr)
		}
	}
}

func (k *KafkaReader) Close() error {
	return k.reader.Close()
}
