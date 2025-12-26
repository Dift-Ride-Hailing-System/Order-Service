package dto

import (
	"time"

	"dift_backend_go/order-service/internal/model"
	pb "dift_backend_go/order-service/proto/pb"

	"github.com/golang/protobuf/proto"
)


func OrderMatchingToPayload(
	m model.OrderMatching,
) ([]byte, error) {
	return proto.Marshal(OrderMatchingToPB(m))
}

// --------------------------
// PB - Model


// OrderMatchingFromPB แปลง pb - model
func OrderMatchingFromPB(pbReq *pb.OrderMatchingRequest) model.OrderMatching {
	return model.OrderMatching{
		OrderID:         pbReq.OrderId,
		UserID:          pbReq.UserId,
		CarType:         pbReq.CarType,
		PickupLocation:  pbReq.PickupLocation,
		DropoffLocation: pbReq.DropoffLocation,
		FinalTotal:      pbReq.FinalTotal,
		CouponCode:      pbReq.CouponCode,
		Distance:        pbReq.Distance,
		Duration:        pbReq.Duration,
		PickupPolyline:  pbReq.PickupPolyline,
		DropoffPolyline: pbReq.DropoffPolyline,
		Timestamp:       time.Unix(pbReq.Timestamp, 0),
	}
}

// Model - PB


// OrderMatchingToPB แปลง model - pb
func OrderMatchingToPB(m model.OrderMatching) *pb.OrderMatchingRequest {
	return &pb.OrderMatchingRequest{
		OrderId:         m.OrderID,
		UserId:          m.UserID,
		CarType:         m.CarType,
		PickupLocation:  m.PickupLocation,
		DropoffLocation: m.DropoffLocation,
		FinalTotal:      m.FinalTotal,
		CouponCode:      m.CouponCode,
		Distance:        m.Distance,
		Duration:        m.Duration,
		PickupPolyline:  m.PickupPolyline,
		DropoffPolyline: m.DropoffPolyline,
		Timestamp:       m.Timestamp.Unix(),
	}
}
