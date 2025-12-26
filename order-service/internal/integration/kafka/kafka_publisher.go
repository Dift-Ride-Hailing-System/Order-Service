package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer interface {
	SendMessage(topic, key string, value []byte) error
	Close() error
}

type KafkaWriter struct {
	writer *kafka.Writer
}

func NewKafkaWriter(brokers []string) *KafkaWriter {
	return &KafkaWriter{
		writer: &kafka.Writer{
			Addr:         kafka.TCP(brokers...),
			Balancer:     &kafka.LeastBytes{},
			RequiredAcks: kafka.RequireOne, // ไว + ปลอดภัยระดับกลาง
		},
	}
}

func (k *KafkaWriter) SendMessage(topic, key string, value []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	msg := kafka.Message{
		Topic: topic,
		Key:   []byte(key),
		Value: value,
		Time:  time.Now(),
	}

	err := k.writer.WriteMessages(ctx, msg)
	if err != nil {
		log.Println("[KafkaProducer] Failed:", err)
		return err
	}

	// log.Println("[KafkaProducer] Published →", topic, "key:", key)
	return nil
}

func (k *KafkaWriter) Close() error {
	return k.writer.Close()
}
