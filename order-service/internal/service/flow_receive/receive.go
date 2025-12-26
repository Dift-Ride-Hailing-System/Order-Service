package flow_receive

import (
	"context"
	"dift_backend_go/order-service/internal/model"
	"dift_backend_go/order-service/internal/observability/logger"
	"dift_backend_go/order-service/internal/service/core"
	"encoding/json"
	"errors"
	"time"
)

/*****************use case *****************************/

type ReceiveFlow struct {
	orderSvc *core.OrderService
}

func NewReceiveFlow(orderSvc *core.OrderService) *ReceiveFlow {
	return &ReceiveFlow{orderSvc: orderSvc}
}

// ---------- Receive Travel Order ---------- /

func (f *ReceiveFlow) ReceiveTravelOrder(
	ctx context.Context,
	req model.TravelRequest,
) (*model.Order, error) {

	// 1) Validate (domain-level)
	if req.RouteID == "" ||
		req.UserID == "" ||
		req.Pickup.Address == "" ||
		req.Dropoff.Address == "" {
		return nil, errors.New("invalid travel request")
	}

	// 2) Idempotency
	lockKey := "travel:" + req.RouteID
	if !f.orderSvc.Idem().TryLock(lockKey) {
		logger.Info("ReceiveTravelOrder: duplicate ignored route=%s", req.RouteID)
		return nil, nil
	}
	defer f.orderSvc.Idem().Unlock(lockKey)

	// 3) Build Order
	now := time.Now()

	order := &model.Order{
		OrderID:   req.RouteID,
		UserID:    req.UserID,
		Pickup:    req.Pickup.Address,
		Dropoff:   req.Dropoff.Address,
		Status:    string(model.StatusWaitingCoupon),
		Estimated: req.Price,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// 4) Store in cache
	data, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	ctxSet, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	if err := f.orderSvc.Cache().Set(
		ctxSet,
		"order:"+order.OrderID,
		data,
		int(f.orderSvc.CacheTTL().Seconds()),
	); err != nil {
		return nil, err
	}

	logger.Info(
		"ReceiveTravelOrder: order cached orderID=%s userID=%s",
		order.OrderID,
		order.UserID,
	)

	return order, nil
}
