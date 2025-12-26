package adapter

import (
	"context"

	port "dift_backend_go/order-service/internal/interface"
	"dift_backend_go/order-service/internal/model"
	"dift_backend_go/order-service/internal/observability/logger"
	"dift_backend_go/order-service/internal/service/flow_cancel"
)

// DriverCancelConsumer เป็น consumer adapter สำหรับ driver cancel
// ทำหน้าที่รับ model แล้วส่งเข้า business flow เท่านั้น
type DriverCancelConsumer struct {
	flow *flow_cancel.CancelFlow
}

// compile-time check ว่า implement port จริง
var _ port.DriverCancelConsumerPort = (*DriverCancelConsumer)(nil)

// NewDriverCancelConsumer constructor
func NewDriverCancelConsumer(flow *flow_cancel.CancelFlow) *DriverCancelConsumer {
	return &DriverCancelConsumer{flow: flow}
}

// Handle รับ model.DriverCancelResponse ตาม contract ของ port
func (c *DriverCancelConsumer) Handle(event model.DriverCancelResponse) error {
	// สร้าง context สำหรับ flow
	ctx := context.Background()

	// ส่งเข้า business flow
	if err := c.flow.HandleDriverCancel(ctx, event); err != nil {
		logger.Warn(
			"DriverCancelConsumer: failed order=%s err=%v",
			event.OrderID,
			err,
		)
		return err
	}

	logger.Info(
		"DriverCancelConsumer: handled driver cancel order=%s status=%s",
		event.OrderID,
		event.Status,
	)

	return nil
}
