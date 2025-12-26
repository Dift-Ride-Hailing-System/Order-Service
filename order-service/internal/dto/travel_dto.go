package dto

import (
	"time"

	"dift_backend_go/order-service/internal/model"
	pb "dift_backend_go/order-service/proto/pb"

	"github.com/golang/protobuf/proto"
)

//
// ==========================
// Travel Request DTO
// ==========================
//

// --------------------------
// Payload → Model
// --------------------------

// TravelRequestFromPayload
// ใช้ใน Consumer
// - payload (Kafka)
// - proto.Unmarshal → pb
// - pb → model
func TravelRequestFromPayload(
	payload []byte,
) (model.TravelRequest, error) {
	var pbReq pb.TravelRequest
	if err := proto.Unmarshal(payload, &pbReq); err != nil {
		return model.TravelRequest{}, err
	}

	return TravelRequestFromPB(&pbReq), nil
}

// --------------------------
// PB → Model
// --------------------------

func TravelLocationFromPB(pbLoc *pb.TravelLocation) model.TravelLocation {
	return model.TravelLocation{
		Lat:     pbLoc.Lat,
		Lng:     pbLoc.Lng,
		Address: pbLoc.Address,
	}
}

func TravelRequestFromPB(pbMsg *pb.TravelRequest) model.TravelRequest {
	return model.TravelRequest{
		RouteID:   pbMsg.RouteId,
		UserID:    pbMsg.UserId,
		Pickup:    TravelLocationFromPB(pbMsg.Pickup),
		Dropoff:   TravelLocationFromPB(pbMsg.Dropoff),
		CarType:   pbMsg.CarType,
		Distance:  pbMsg.Distance,
		Duration:  pbMsg.Duration,
		Price:     pbMsg.Price,
		Currency:  pbMsg.Currency,
		Timestamp: time.Unix(pbMsg.Timestamp, 0),
	}
}

// --------------------------
// Model → PB
// --------------------------

func TravelLocationToPB(loc model.TravelLocation) *pb.TravelLocation {
	return &pb.TravelLocation{
		Lat:     loc.Lat,
		Lng:     loc.Lng,
		Address: loc.Address,
	}
}

func TravelRequestToPB(m model.TravelRequest) *pb.TravelRequest {
	return &pb.TravelRequest{
		RouteId:   m.RouteID,
		UserId:    m.UserID,
		Pickup:    TravelLocationToPB(m.Pickup),
		Dropoff:   TravelLocationToPB(m.Dropoff),
		CarType:   m.CarType,
		Distance:  m.Distance,
		Duration:  m.Duration,
		Price:     m.Price,
		Currency:  m.Currency,
		Timestamp: m.Timestamp.Unix(),
	}
}
