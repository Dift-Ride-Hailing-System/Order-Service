package flow_cancel

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"dift_backend_go/order-service/internal/model"
	"dift_backend_go/order-service/internal/observability/logger"
)

/* ---------- Passenger Cancel ---------- */

// HandlePassengerCancel
// business logic เมื่อ passenger ยกเลิกงาน
func (f *CancelFlow) HandlePassengerCancel(
	ctx context.Context,
	req model.PassengerCancelRequest,
) error {

	// 1) Idempotency lock
	key := "cancel:" + req.OrderID
	if !f.orderSvc.Idem().TryLock(key) {
		return errors.New("duplicate cancel request")
	}
	defer f.orderSvc.Idem().Unlock(key)

	// 2) Context timeout
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	// 3) Load order from cache
	data, err := f.orderSvc.Cache().Get(ctx, "order:"+req.OrderID)
	if err != nil || data == nil {
		return errors.New("order not found")
	}

	var order model.Order
	if err := json.Unmarshal(data, &order); err != nil {
		return err
	}

	// 4) Validate status transition
	if err := model.ValidateOrderStatusTransition(
		model.OrderStatus(order.Status),
		model.StatusCancelled,
	); err != nil && order.Status != string(model.StatusWaitingCoupon) {
		return errors.New("cannot cancel order in current status")
	}

	// 5) Update status
	order.Status = string(model.StatusCancelled)
	order.UpdatedAt = time.Now()

	updated, _ := json.Marshal(order)
	if err := f.orderSvc.Cache().Set(
		ctx,
		"order:"+order.OrderID,
		updated,
		int(f.orderSvc.CacheTTL().Seconds()),
	); err != nil {
		logger.Error(
			"HandlePassengerCancel: failed to update order cache order=%s err=%v",
			order.OrderID,
			err,
		)
		return err
	}

	// 6) Notify driver (business decision)
	if order.DriverID != "" {
		f.notifyPassengerCancelToDriver(model.PassengerCancelRequest{
			OrderID:   order.OrderID,
			UserID:    req.UserID,
			Timestamp: time.Now(),
		})
	}

	logger.Info("HandlePassengerCancel: passenger cancelled order=%s", order.OrderID)
	return nil
}

// notifyPassengerCancelToDriver เป็น method private ของ CancelFlow
// แยก logic ส่ง notification ไป driver
func (f *CancelFlow) notifyPassengerCancelToDriver(req model.PassengerCancelRequest) {
	// TODO: implement actual notification logic, เช่น Kafka producer หรือ push
	logger.Info("Notify driver: order=%s user=%s", req.OrderID, req.UserID)
}
