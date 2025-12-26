package grpc

import (
	"context"
	"time"

	"dift_backend_go/order-service/internal/dto"
	gprc "dift_backend_go/order-service/internal/interface/gprc"
	"dift_backend_go/order-service/internal/model"
	pb "dift_backend_go/order-service/proto/pb/usercoupon"

	"google.golang.org/grpc"
)

const defaultCallTimeout = 2 * time.Second

type UserCouponAdapter struct {
	client pb.UserCouponServiceClient
	conn   *grpc.ClientConn
}

// compile-time check
var _ gprc.UserCouponService = (*UserCouponAdapter)(nil)

func NewUserCouponAdapter(
	conn *grpc.ClientConn,
) *UserCouponAdapter {
	return &UserCouponAdapter{
		client: pb.NewUserCouponServiceClient(conn),
		conn:   conn,
	}
}

func (a *UserCouponAdapter) Close() error {
	if a.conn == nil {
		return nil
	}
	return a.conn.Close()
}

func (a *UserCouponAdapter) ListCoupons(
	parentCtx context.Context,
	userID string,
) (model.ListUserCouponsResponse, error) {

	ctx, cancel := context.WithTimeout(parentCtx, defaultCallTimeout)
	defer cancel()

	resp, err := a.client.ListUserCoupons(
		ctx,
		&pb.ListUserCouponsRequest{UserId: userID},
	)
	if err != nil {
		return model.ListUserCouponsResponse{}, err
	}

	return dto.ListUserCouponsResponseFromPB(resp), nil
}
