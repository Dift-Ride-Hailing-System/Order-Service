package flow_confirm

import (
	"context"
	"dift_backend_go/order-service/internal/model"
	"dift_backend_go/order-service/internal/observability/logger"
	"dift_backend_go/order-service/internal/service/core"
	"encoding/json"
	"errors"
	"time"
)

const (
	confirmTimeout = 2 * time.Second
)

/* ---------- use case ---------- */

type ConfirmFlow struct {
	orderSvc *core.OrderService
}

func NewConfirmFlow(
	orderSvc *core.OrderService,
) *ConfirmFlow {
	return &ConfirmFlow{
		orderSvc: orderSvc,
	}
}

/* ---------- Confirm Order ---------- */

func (f *ConfirmFlow) ConfirmOrder(
	ctx context.Context,
	orderID string,
	couponCode string,
	paymentMethod string, // reserved for future
) (*model.OrderMatching, error) {

	// 1) Idempotency
	lockKey := "confirm:" + orderID
	if !f.orderSvc.Idem().TryLock(lockKey) {
		return nil, errors.New("duplicate confirm request")
	}
	defer f.orderSvc.Idem().Unlock(lockKey)

	// 2) Load order from cache
	ctxGet, cancel := context.WithTimeout(ctx, confirmTimeout)
	defer cancel()

	order, err := f.loadOrder(ctxGet, orderID)
	if err != nil {
		return nil, err
	}

	// 3) Validate status transition
	if err := model.ValidateOrderStatusTransition(
		model.OrderStatus(order.Status),
		model.StatusMatchingSent,
	); err != nil {
		return nil, err
	}

	// 4) Coupon calculation (domain orchestration)
	if couponCode != "" {
		if err := f.applyCoupon(ctx, order, couponCode); err != nil {
			return nil, err
		}
	} else {
		order.FinalTotal = order.Estimated
		order.Discount = 0
		order.CouponCode = ""
	}

	// 5) Update status
	order.Status = string(model.StatusMatchingSent)
	order.UpdatedAt = time.Now()

	// 6) Persist order
	if err := f.saveOrder(ctx, order); err != nil {
		return nil, err
	}

	// 7) Build matching payload
	return f.buildMatching(order), nil
}

/* ---------- helpers ---------- */

func (f *ConfirmFlow) loadOrder(
	ctx context.Context,
	orderID string,
) (*model.Order, error) {

	data, err := f.orderSvc.Cache().Get(ctx, "order:"+orderID)
	if err != nil || data == nil {
		return nil, errors.New("order not found")
	}

	var order model.Order
	if err := json.Unmarshal(data, &order); err != nil {
		logger.Error(
			"ConfirmFlow: unmarshal order=%s err=%v",
			orderID,
			err,
		)
		return nil, err
	}

	return &order, nil
}

func (f *ConfirmFlow) saveOrder(
	ctx context.Context,
	order *model.Order,
) error {

	data, _ := json.Marshal(order)

	ctxSet, cancel := context.WithTimeout(ctx, confirmTimeout)
	defer cancel()

	return f.orderSvc.Cache().Set(
		ctxSet,
		"order:"+order.OrderID,
		data,
		int(f.orderSvc.CacheTTL().Seconds()),
	)
}

func (f *ConfirmFlow) applyCoupon(
	ctx context.Context,
	order *model.Order,
	couponCode string,
) error {

	// ensure coupon components exist
	if f.orderSvc.CouponFetcher() == nil ||
		f.orderSvc.CouponPrecalc() == nil ||
		f.orderSvc.CouponApplier() == nil {
		return errors.New("coupon service not available")
	}

	// 1) fetch coupons
	couponResp, err :=
		f.orderSvc.CouponFetcher().GetCoupons(ctx, order.UserID)
	if err != nil {
		return err
	}

	// 2) precompute prices
	if err := f.orderSvc.CouponPrecalc().PrecomputePrices(
		ctx,
		order.UserID,
		order.Estimated,
		extractCouponCodes(couponResp),
	); err != nil {
		return err
	}

	// 3) apply selected coupon
	price, err :=
		f.orderSvc.CouponApplier().ApplyPrecomputedCoupon(
			ctx,
			order.UserID,
			couponCode,
		)
	if err != nil {
		return err
	}

	order.FinalTotal = price.FinalTotal
	order.Discount = price.Discount
	order.CouponCode = couponCode

	return nil
}

func (f *ConfirmFlow) buildMatching(
	order *model.Order,
) *model.OrderMatching {

	return &model.OrderMatching{
		OrderID:         order.OrderID,
		UserID:          order.UserID,
		CarType:         "",
		PickupLocation:  order.Pickup,
		DropoffLocation: order.Dropoff,
		FinalTotal:      order.FinalTotal,
		CouponCode:      order.CouponCode,
		Distance:        0,
		Duration:        0,
		PickupPolyline:  "",
		DropoffPolyline: "",
		Timestamp:       time.Now(),
	}
}

/* ---------- util ---------- */

func extractCouponCodes(
	resp model.ListUserCouponsResponse,
) []string {

	codes := make([]string, 0, len(resp.Coupons))
	for _, c := range resp.Coupons {
		codes = append(codes, c.Code)
	}
	return codes
}

/* ---------- Get Order ---------- */

func (f *ConfirmFlow) GetOrder(
	ctx context.Context,
	orderID string,
) (*model.Order, error) {

	ctxGet, cancel := context.WithTimeout(ctx, confirmTimeout)
	defer cancel()

	return f.loadOrder(ctxGet, orderID)
}
