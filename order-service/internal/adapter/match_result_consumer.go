package adapter

import (
	"context"

	port "dift_backend_go/order-service/internal/interface"
	"dift_backend_go/order-service/internal/model"
	"dift_backend_go/order-service/internal/observability/logger"
	"dift_backend_go/order-service/internal/service/core"
	"dift_backend_go/order-service/internal/service/flow_match"
)

// MatchResultConsumer เป็น adapter สำหรับรับผล match
// ทำหน้าที่รับ model แล้วส่งเข้า business flow
type MatchResultConsumer struct {
	orderSvc *core.OrderService
}

// compile-time check ว่า implement port จริง
var _ port.MatchResultHandler = (*MatchResultConsumer)(nil)

// NewMatchResultConsumer constructor
func NewMatchResultConsumer(
	orderSvc *core.OrderService,
) *MatchResultConsumer {
	return &MatchResultConsumer{orderSvc: orderSvc}
}

// Handle รับ model.OrderMatchNotification ตาม contract ของ port
func (c *MatchResultConsumer) Handle(
	match model.OrderMatchNotification,
) error {

	// business logic
	if err := flow_match.ApplyMatchResult(
		context.Background(),
		c.orderSvc,
		match,
	); err != nil {
		logger.Warn(
			"MatchResultConsumer: failed orderID=%s err=%v",
			match.OrderID,
			err,
		)
		return err
	}

	return nil
}
