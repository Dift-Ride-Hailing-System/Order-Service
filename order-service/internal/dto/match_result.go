package dto

import (
	"time"

	"dift_backend_go/order-service/internal/model"
	pb "dift_backend_go/order-service/proto/pb"

	"github.com/golang/protobuf/proto"
)

//
// ==========================
// Match Result Notification DTO
// ==========================
//
// ความรับผิดชอบของไฟล์นี้:
// - เป็น boundary ระหว่าง transport (Kafka / Protobuf) กับ domain (model)
// - เป็นที่เดียวที่รู้จัก protobuf schema
// - เป็นที่เดียวที่ marshal / unmarshal proto
// - แปลงข้อมูล:
//     payload ↔ pb ↔ model
//

// ==================================================
// Payload → Model
// ==================================================

// NotificationFromPayload
// ใช้ใน Consumer
// - payload (Kafka)
// - proto.Unmarshal → pb
// - pb → model
func NotificationFromPayload(
	payload []byte,
) (model.OrderMatchNotification, error) {
	var pbMsg pb.OrderMatchNotification
	if err := proto.Unmarshal(payload, &pbMsg); err != nil {
		return model.OrderMatchNotification{}, err
	}

	return NotificationFromPB(&pbMsg), nil
}

// ==================================================
// PB → Model
// ==================================================

// NotificationFromPB
// แปลง protobuf → domain model
func NotificationFromPB(
	pbMsg *pb.OrderMatchNotification,
) model.OrderMatchNotification {
	return model.OrderMatchNotification{
		OrderID:                 pbMsg.OrderId,
		Status:                  pbMsg.Status,
		DriverID:                pbMsg.DriverId,
		DriverName:              pbMsg.DriverName,
		DriverCarModel:          pbMsg.DriverCarModel,
		DriverAvatarURL:         pbMsg.DriverAvatarUrl,
		CarPlate:                pbMsg.CarPlate,
		CarType:                 pbMsg.CarType,
		DriverLat:               pbMsg.DriverLat,
		DriverLng:               pbMsg.DriverLng,
		PickupLat:               pbMsg.PickupLat,
		PickupLng:               pbMsg.PickupLng,
		PickupAddress:           pbMsg.PickupAddress,
		DropoffLat:              pbMsg.DropoffLat,
		DropoffLng:              pbMsg.DropoffLng,
		DropoffAddress:          pbMsg.DropoffAddress,
		DistancePickupToDropoff: pbMsg.DistancePickupToDropoff,
		DurationTotal:           pbMsg.DurationTotal,
		RoutePolyline:           pbMsg.RoutePolyline,
		Price:                   pbMsg.Price,
		Timestamp:               time.Unix(pbMsg.Timestamp, 0),
	}
}

// ==================================================
// Model → PB
// ==================================================

// NotificationToPB
// ใช้กรณีที่ระบบนี้ต้องส่ง MatchResult notification ออกไป (future-proof)
func NotificationToPB(
	m model.OrderMatchNotification,
) *pb.OrderMatchNotification {
	ts := m.Timestamp
	if ts.IsZero() {
		ts = time.Now()
	}

	return &pb.OrderMatchNotification{
		OrderId:                 m.OrderID,
		Status:                  m.Status,
		DriverId:                m.DriverID,
		DriverName:              m.DriverName,
		DriverCarModel:          m.DriverCarModel,
		DriverAvatarUrl:         m.DriverAvatarURL,
		CarPlate:                m.CarPlate,
		CarType:                 m.CarType,
		DriverLat:               m.DriverLat,
		DriverLng:               m.DriverLng,
		PickupLat:               m.PickupLat,
		PickupLng:               m.PickupLng,
		PickupAddress:           m.PickupAddress,
		DropoffLat:              m.DropoffLat,
		DropoffLng:              m.DropoffLng,
		DropoffAddress:          m.DropoffAddress,
		DistancePickupToDropoff: m.DistancePickupToDropoff,
		DurationTotal:           m.DurationTotal,
		RoutePolyline:           m.RoutePolyline,
		Price:                   m.Price,
		Timestamp:               ts.Unix(),
	}
}

// ==================================================
// Model → Payload
// ==================================================

// NotificationToPayload
// ใช้ใน Producer (ถ้ามีในอนาคต)
// - model → pb
// - proto.Marshal → payload
func NotificationToPayload(
	m model.OrderMatchNotification,
) ([]byte, error) {
	pbMsg := NotificationToPB(m)
	return proto.Marshal(pbMsg)
}
