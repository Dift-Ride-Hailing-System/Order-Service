package coupon

import (
	"context"
	"encoding/json"
	"time"

	port "dift_backend_go/order-service/internal/interface"
	gprc "dift_backend_go/order-service/internal/interface/gprc"
	"dift_backend_go/order-service/internal/model"
)

const userCouponTimeout = 3 * time.Second

type UserCouponFetcher struct {
	service  gprc.UserCouponService
	cache    port.Cache
	cacheTTL time.Duration
}

func NewUserCouponFetcher(
	service gprc.UserCouponService,
	cache port.Cache,
	ttl time.Duration,
) *UserCouponFetcher {
	return &UserCouponFetcher{
		service:  service,
		cache:    cache,
		cacheTTL: ttl,
	}
}

func (f *UserCouponFetcher) GetCoupons(
	ctx context.Context,
	userID string,
) (model.ListUserCouponsResponse, error) {

	ctx, cancel := context.WithTimeout(ctx, userCouponTimeout)
	defer cancel()

	key := "user_coupons:" + userID

	if data, _ := f.cache.Get(ctx, key); data != nil {
		var cached model.ListUserCouponsResponse
		if err := json.Unmarshal(data, &cached); err == nil {
			return cached, nil
		}
	}

	resp, err := f.service.ListCoupons(ctx, userID)
	if err != nil {
		return model.ListUserCouponsResponse{}, err
	}

	data, _ := json.Marshal(resp)
	_ = f.cache.Set(ctx, key, data, int(f.cacheTTL.Seconds()))

	return resp, nil
}
