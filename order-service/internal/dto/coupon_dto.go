package dto

import (
	"dift_backend_go/order-service/internal/model"
	pb "dift_backend_go/order-service/proto/pb/coupon"
)

// Model → PB
func ApplyCouponInputToPB(
	input model.ApplyCouponInput,
) *pb.ApplyCouponRequest {
	return &pb.ApplyCouponRequest{
		UserId:     input.UserID,
		CouponCode: input.CouponCode,
		OrderTotal: input.OrderTotal,
	}
}

// PB → Model
func ApplyCouponResultFromPB(
	resp *pb.ApplyCouponResponse,
) model.ApplyCouponResult {
	return model.ApplyCouponResult{
		FinalTotal: resp.FinalTotal,
		Discount:   resp.Discount,
		Valid:      resp.Valid,
		Message:    resp.Message,
	}
}
