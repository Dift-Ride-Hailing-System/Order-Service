package dto

import (
	"time"

	"dift_backend_go/order-service/internal/model"
	pb "dift_backend_go/order-service/proto/pb"

	"github.com/golang/protobuf/proto"
)

//
// ==========================
// Trip History DTO
//********************************************************
// Payload - Model
// ==================================================

// TripHistoryEventFromPayload
// ใช้ใน Consumer (Kafka / Redpanda)
func TripHistoryEventFromPayload(
	payload []byte,
) (model.TripHistoryEvent, error) {

	var pbMsg pb.TripHistoryEvent
	if err := proto.Unmarshal(payload, &pbMsg); err != nil {
		return model.TripHistoryEvent{}, err
	}

	return TripHistoryEventFromPB(&pbMsg), nil
}

// PB - Model

// TripHistoryEventFromPB
// แปลง protobuf -domain model
func TripHistoryEventFromPB(
	pbMsg *pb.TripHistoryEvent,
) model.TripHistoryEvent {

	var ts time.Time
	if pbMsg.Timestamp > 0 {
		ts = time.Unix(pbMsg.Timestamp, 0)
	}

	return model.TripHistoryEvent{
		OrderID:         pbMsg.OrderId,
		UserID:          pbMsg.UserId,
		Status:          pbMsg.Status,
		DriverID:        pbMsg.DriverId,
		DriverName:      pbMsg.DriverName,
		DriverCarModel:  pbMsg.DriverCarModel,
		DriverAvatarURL: pbMsg.DriverAvatarUrl,
		CarPlate:        pbMsg.CarPlate,
		CarType:         pbMsg.CarType,
		PickupLocation:  pbMsg.PickupLocation,
		DropoffLocation: pbMsg.DropoffLocation,
		Distance:        pbMsg.Distance,
		Duration:        pbMsg.Duration,
		PickupPolyline:  pbMsg.PickupPolyline,
		DropoffPolyline: pbMsg.DropoffPolyline,
		FinalTotal:      pbMsg.FinalTotal,
		CouponCode:      pbMsg.CouponCode,
		Timestamp:       ts,
		Metadata:        pbMsg.Metadata,
	}
}

// Model - Payload

// TripHistoryEventToPayload
// ใช้ใน Producer
func TripHistoryEventToPayload(
	m model.TripHistoryEvent,
) ([]byte, error) {

	pbMsg := TripHistoryEventToPB(m)
	return proto.Marshal(pbMsg)
}

// Model - PB

// TripHistoryEventToPB
// แปลง domain model -protobuf
func TripHistoryEventToPB(
	m model.TripHistoryEvent,
) *pb.TripHistoryEvent {

	ts := m.Timestamp
	if ts.IsZero() {
		ts = time.Now()
	}

	metadata := m.Metadata
	if metadata == nil {
		metadata = map[string]string{}
	}

	return &pb.TripHistoryEvent{
		OrderId:         m.OrderID,
		UserId:          m.UserID,
		Status:          m.Status,
		DriverId:        m.DriverID,
		DriverName:      m.DriverName,
		DriverCarModel:  m.DriverCarModel,
		DriverAvatarUrl: m.DriverAvatarURL,
		CarPlate:        m.CarPlate,
		CarType:         m.CarType,
		PickupLocation:  m.PickupLocation,
		DropoffLocation: m.DropoffLocation,
		Distance:        m.Distance,
		Duration:        m.Duration,
		PickupPolyline:  m.PickupPolyline,
		DropoffPolyline: m.DropoffPolyline,
		FinalTotal:      m.FinalTotal,
		CouponCode:      m.CouponCode,
		Timestamp:       ts.Unix(),
		Metadata:        metadata,
	}
}
