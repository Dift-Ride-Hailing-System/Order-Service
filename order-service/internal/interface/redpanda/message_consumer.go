package redpanda

// MessageConsumer
// - abstraction ของ broker consumer
type MessageConsumer interface {
	Start(handler func(key string, payload []byte) error) error
	Close() error
}
