package redpanda

// MessageProducer
// - abstraction ของ broker producer

type MessageProducer interface {
	Send(topic, key string, payload []byte) error
	Close() error
}
