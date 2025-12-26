package grpcinfra

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialCouponService(
	address string,
	timeout time.Duration,
) (*grpc.ClientConn, error) {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return grpc.DialContext(
		ctx,
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
}
