package adapter

import (
	"context"

	port "dift_backend_go/order-service/internal/interface"
	"dift_backend_go/order-service/internal/model"
	"dift_backend_go/order-service/internal/observability/logger"
	"dift_backend_go/order-service/internal/service/flow_receive"
)

// ReceiveTravelConsumer เป็น adapter สำหรับรับ TravelRequest
// implements port.TravelRequestConsumer
type ReceiveTravelConsumer struct {
	flow *flow_receive.ReceiveFlow
}

// compile-time check
var _ port.TravelRequestConsumer = (*ReceiveTravelConsumer)(nil)

// NewReceiveTravelConsumer constructor
func NewReceiveTravelConsumer(
	flow *flow_receive.ReceiveFlow,
) *ReceiveTravelConsumer {
	return &ReceiveTravelConsumer{flow: flow}
}

// Handle รับ model.TravelRequest แล้วส่งเข้า business flow
func (c *ReceiveTravelConsumer) Handle(
	req model.TravelRequest,
) (*model.Order, error) {

	order, err := c.flow.ReceiveTravelOrder(
		context.Background(),
		req,
	)
	if err != nil {
		logger.Warn(
			"ReceiveTravelConsumer: failed route=%s err=%v",
			req.RouteID,
			err,
		)
		return nil, err
	}

	return order, nil
}
