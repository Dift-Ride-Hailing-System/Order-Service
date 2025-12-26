package flow_cancel

import (
	"context"
	"encoding/json"
	"time"

	"dift_backend_go/order-service/internal/model"
	"dift_backend_go/order-service/internal/observability/logger"
	"dift_backend_go/order-service/internal/service/core"
)

/* ---------- use case ---------- */

type CancelFlow struct {
	orderSvc *core.OrderService
}

func NewCancelFlow(orderSvc *core.OrderService) *CancelFlow {
	return &CancelFlow{orderSvc: orderSvc}
}

/* ---------- Driver Cancel ---------- */

// HandleDriverCancel
// business logic เมื่อ driver ยกเลิกงาน
func (f *CancelFlow) HandleDriverCancel(
	ctx context.Context,
	event model.DriverCancelResponse,
) error {

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	data, err := f.orderSvc.Cache().Get(ctx, "order:"+event.OrderID)
	if err != nil || data == nil {
		logger.Warn(
			"HandleDriverCancel: order not found orderID=%s",
			event.OrderID,
		)
		return nil
	}

	var order model.Order
	if err := json.Unmarshal(data, &order); err != nil {
		return err
	}

	// validate status transition
	if err := model.ValidateOrderStatusTransition(
		model.OrderStatus(order.Status),
		model.OrderStatus(event.Status),
	); err != nil {
		logger.Warn(
			"HandleDriverCancel: invalid transition order=%s current=%s next=%s",
			order.OrderID,
			order.Status,
			string(event.Status),
		)
		return nil
	}

	// skip if no change
	if order.Status == string(event.Status) {
		return nil
	}

	order.Status = string(event.Status) // แปลง type ให้ตรงกับ string
	order.UpdatedAt = time.Now()

	updated, _ := json.Marshal(order)
	if err := f.orderSvc.Cache().Set(
		ctx,
		"order:"+order.OrderID,
		updated,
		int(f.orderSvc.CacheTTL().Seconds()),
	); err != nil {
		logger.Error(
			"HandleDriverCancel: failed to update order cache order=%s err=%v",
			order.OrderID,
			err,
		)
		return err
	}

	logger.Info(
		"HandleDriverCancel: updated order=%s status=%s",
		order.OrderID,
		string(event.Status),
	)

	return nil
}
