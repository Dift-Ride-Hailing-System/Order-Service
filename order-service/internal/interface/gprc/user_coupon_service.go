package gprc_port

import (
	"context"

	"dift_backend_go/order-service/internal/model"
)

type UserCouponService interface {
	ListCoupons(
		ctx context.Context,
		userID string,
	) (model.ListUserCouponsResponse, error)
}
