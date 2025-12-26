package gprc_port

import (
	"context"

	"dift_backend_go/order-service/internal/model"
)

type CouponService interface {
	ApplyCoupon(
		ctx context.Context,
		input model.ApplyCouponInput,
	) (model.ApplyCouponResult, error)
}
