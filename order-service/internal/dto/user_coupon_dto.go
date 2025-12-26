package dto

import (
	"time"

	"dift_backend_go/order-service/internal/model"
	pb "dift_backend_go/order-service/proto/pb/usercoupon"
)

func UserCouponFromPB(pbCoupon *pb.UserCoupon) model.UserCoupon {
	return model.UserCoupon{
		Code:        pbCoupon.Code,
		Title:       pbCoupon.Title,
		Description: pbCoupon.Description,
		Discount:    pbCoupon.Discount,
		Currency:    pbCoupon.Currency,
		ValidUntil:  time.Unix(pbCoupon.ValidUntil, 0),
	}
}

func ListUserCouponsResponseFromPB(
	pbResp *pb.ListUserCouponsResponse,
) model.ListUserCouponsResponse {

	coupons := make([]model.UserCoupon, len(pbResp.Coupons))
	for i, c := range pbResp.Coupons {
		coupons[i] = UserCouponFromPB(c)
	}

	return model.ListUserCouponsResponse{
		Coupons: coupons,
	}
}
