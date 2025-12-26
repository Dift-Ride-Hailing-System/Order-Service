package grpc

import (
	"context"

	"dift_backend_go/order-service/internal/dto"
	gprc "dift_backend_go/order-service/internal/interface/gprc"
	"dift_backend_go/order-service/internal/model"
	pb "dift_backend_go/order-service/proto/pb/coupon"
)

type CouponServiceAdapter struct {
	client pb.CouponServiceClient
}

// compile-time check
var _ gprc.CouponService = (*CouponServiceAdapter)(nil)

func NewCouponServiceAdapter(
	client pb.CouponServiceClient,
) *CouponServiceAdapter {
	return &CouponServiceAdapter{client: client}
}

func (a *CouponServiceAdapter) ApplyCoupon(
	ctx context.Context,
	input model.ApplyCouponInput,
) (model.ApplyCouponResult, error) {

	resp, err := a.client.ApplyCoupon(
		ctx,
		dto.ApplyCouponInputToPB(input),
	)
	if err != nil {
		return model.ApplyCouponResult{}, err
	}

	return dto.ApplyCouponResultFromPB(resp), nil
}
