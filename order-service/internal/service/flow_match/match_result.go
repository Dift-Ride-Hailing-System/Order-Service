package flow_match

import (
	"context"
	"encoding/json"
	"time"

	"dift_backend_go/order-service/internal/model"
	"dift_backend_go/order-service/internal/observability/logger"
	"dift_backend_go/order-service/internal/service/core"
)

// ApplyMatchResult
// business logic สำหรับ apply ผลการจับคู่
func ApplyMatchResult(
	ctx context.Context,
	orderSvc *core.OrderService,
	match model.OrderMatchNotification,
) error {

	// 1) get order from cache
	ctxGet, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	data, err := orderSvc.Cache().Get(ctxGet, "order:"+match.OrderID)
	if err != nil || data == nil {
		logger.Warn(
			"ApplyMatchResult: order not found orderID=%s",
			match.OrderID,
		)
		return nil
	}

	var order model.Order
	if err := json.Unmarshal(data, &order); err != nil {
		return err
	}

	// 2) idempotent check
	if order.MatchResult != nil &&
		order.MatchResult.Status == match.Status {
		return nil
	}

	// 3) validate domain transition
	if err := model.ValidateOrderStatusTransition(
		model.OrderStatus(order.Status),
		model.OrderStatus(match.Status),
	); err != nil {
		logger.Warn(
			"ApplyMatchResult: invalid transition orderID=%s from=%s to=%s",
			order.OrderID,
			order.Status,
			match.Status,
		)
		return nil
	}

	// ---------- APPLY MATCH RESULT

	order.Status = match.Status
	order.DriverID = match.DriverID
	order.DriverName = match.DriverName
	order.DriverCar = match.CarPlate
	order.MatchResult = &model.MatchResult{
		Status:   match.Status,
		DriverID: match.DriverID,
	}
	order.UpdatedAt = time.Now()

	// 4) update cache
	updated, err := json.Marshal(order)
	if err != nil {
		return err
	}

	ctxSet, cancelSet := context.WithTimeout(ctx, 2*time.Second)
	defer cancelSet()

	if err := orderSvc.Cache().Set(
		ctxSet,
		"order:"+order.OrderID,
		updated,
		int(orderSvc.CacheTTL().Seconds()),
	); err != nil {
		logger.Error(
			"ApplyMatchResult: cache update failed orderID=%s err=%v",
			order.OrderID,
			err,
		)
		return err
	}

	logger.Info(
		"ApplyMatchResult: order=%s updated status=%s",
		order.OrderID,
		order.Status,
	)

	return nil
}
