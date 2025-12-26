package redpanda

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

type Publisher interface {
	Publish(topic string, key string, value []byte) error
	Close() error
}

type RedpandaPublisher struct {
	writer *kafka.Writer
}

func NewPublisher(brokers []string) *RedpandaPublisher {
	w := &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireOne, // ไวกว่า require all
		Async:        true,             // ทำให้ publisher เบาและเร็วมาก
	}

	return &RedpandaPublisher{writer: w}
}

func (p *RedpandaPublisher) Publish(topic, key string, value []byte) error {
	msg := kafka.Message{
		Topic: topic,
		Key:   []byte(key),
		Value: value,
		Time:  time.Now(),
	}

	return p.writer.WriteMessages(context.Background(), msg)
}

func (p *RedpandaPublisher) Close() error {
	return p.writer.Close()
}
