package adapter

import (
	"dift_backend_go/order-service/internal/dto"
	port "dift_backend_go/order-service/internal/interface"
	"dift_backend_go/order-service/internal/model"
	"dift_backend_go/order-service/internal/observability/logger"
	"dift_backend_go/order-service/internal/service/core"
)

// OrderMatchingProducer เป็น Kafka producer adapter สำหรับ order matching
// implements port.OrderMatchingProducerPort
type OrderMatchingProducer struct {
	service *core.OrderService
	topic   string
}

// compile-time check
var _ port.OrderMatchingProducerPort = (*OrderMatchingProducer)(nil)

// NewOrderMatchingProducer constructor
func NewOrderMatchingProducer(
	service *core.OrderService,
	topic string,
) *OrderMatchingProducer {
	return &OrderMatchingProducer{
		service: service,
		topic:   topic,
	}
}

// Send แปลง model - payload แล้วส่งไป Kafka
func (p *OrderMatchingProducer) Send(matching *model.OrderMatching) error {
	payload, err := dto.OrderMatchingToPayload(*matching)
	if err != nil {
		logger.Error(
			"OrderMatchingProducer: serialize failed order=%s err=%v",
			matching.OrderID,
			err,
		)
		return err
	}

	if err := p.service.Producer().SendMessage(
		p.topic,
		matching.OrderID,
		payload,
	); err != nil {
		logger.Error(
			"OrderMatchingProducer: failed order=%s err=%v",
			matching.OrderID,
			err,
		)
		return err
	}

	logger.Info(
		"OrderMatchingProducer: order=%s sent topic=%s",
		matching.OrderID,
		p.topic,
	)

	return nil
}

// Close ปิด resource
func (p *OrderMatchingProducer) Close() error {
	return p.service.Producer().Close()
}
