package adapter

import (
	"dift_backend_go/order-service/internal/dto"
	port "dift_backend_go/order-service/internal/interface"
	"dift_backend_go/order-service/internal/model"
	"dift_backend_go/order-service/internal/observability/logger"
	"dift_backend_go/order-service/internal/service/core"
)

type PassengerCancelProducer struct {
	service *core.OrderService
	topic   string
}

var _ port.PassengerCancelProducerPort = (*PassengerCancelProducer)(nil)

func NewPassengerCancelProducer(
	service *core.OrderService,
	topic string,
) *PassengerCancelProducer {
	return &PassengerCancelProducer{
		service: service,
		topic:   topic,
	}
}

func (p *PassengerCancelProducer) Send(
	req model.PassengerCancelRequest,
) error {

	// แปลง model - payload
	payload, err := dto.PassengerCancelToPayload(req)
	if err != nil {
		logger.Error(
			"PassengerCancelProducer: serialize failed order=%s err=%v",
			req.OrderID,
			err,
		)
		return err
	}

	// ส่ง payload ไป Kafka
	if err := p.service.Producer().SendMessage(
		p.topic,
		req.OrderID,
		payload,
	); err != nil {
		logger.Error(
			"PassengerCancelProducer: send failed order=%s err=%v",
			req.OrderID,
			err,
		)
		return err
	}

	logger.Info(
		"PassengerCancelProducer: order=%s sent to topic=%s",
		req.OrderID,
		p.topic,
	)

	return nil
}

// Close ปิด resource (optional)
func (p *PassengerCancelProducer) Close() error {
	return p.service.Producer().Close()
}
