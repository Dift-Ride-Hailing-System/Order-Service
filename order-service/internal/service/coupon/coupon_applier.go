package coupon

import (
	"context"
	"encoding/json"
	"errors"

	port "dift_backend_go/order-service/internal/interface"
	"dift_backend_go/order-service/internal/model"
)

type CouponApplier struct {
	cache port.Cache
}

func NewCouponApplier(
	cache port.Cache,
) *CouponApplier {
	return &CouponApplier{
		cache: cache,
	}
}

func (a *CouponApplier) ApplyPrecomputedCoupon(
	ctx context.Context,
	userID string,
	couponCode string,
) (model.CouponPriceResult, error) {

	key := "precomputed_prices:" + userID

	data, err := a.cache.Get(ctx, key)
	if err != nil || data == nil {
		return model.CouponPriceResult{}, errors.New("precomputed prices not found")
	}

	var prices map[string]model.CouponPriceResult
	if err := json.Unmarshal(data, &prices); err != nil {
		return model.CouponPriceResult{}, errors.New("invalid precomputed price format")
	}

	price, ok := prices[couponCode]
	if !ok {
		return model.CouponPriceResult{}, errors.New("coupon not found")
	}

	return price, nil
}
