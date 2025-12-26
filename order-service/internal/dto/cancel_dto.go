package dto

import (
	"time"

	"dift_backend_go/order-service/internal/model"
	pb "dift_backend_go/order-service/proto/pb"

	"github.com/golang/protobuf/proto"
)

//
// ==========================
// Cancel DTO
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
// Passenger Cancel
// ==================================================

// --------------------------
// Model → Payload
// --------------------------

// PassengerCancelToPayload
// ใช้ใน Producer
// - model → pb
// - proto.Marshal → payload
func PassengerCancelToPayload(
	m model.PassengerCancelRequest,
) ([]byte, error) {
	pbReq := PassengerCancelToPB(m)
	return proto.Marshal(pbReq)
}

// --------------------------
// PB → Model
// --------------------------

// PassengerCancelFromPB
// แปลง protobuf → domain model
func PassengerCancelFromPB(
	pbMsg *pb.PassengerCancelRequest,
) model.PassengerCancelRequest {
	return model.PassengerCancelRequest{
		OrderID:   pbMsg.OrderId,
		UserID:    pbMsg.UserId,
		Timestamp: time.Unix(pbMsg.Timestamp, 0),
	}
}

// --------------------------
// Model → PB
// --------------------------

// PassengerCancelToPB
// แปลง domain model → protobuf
func PassengerCancelToPB(
	m model.PassengerCancelRequest,
) *pb.PassengerCancelRequest {
	ts := m.Timestamp
	if ts.IsZero() {
		ts = time.Now()
	}

	return &pb.PassengerCancelRequest{
		OrderId:   m.OrderID,
		UserId:    m.UserID,
		Timestamp: ts.Unix(),
	}
}

// ==================================================
// Driver Cancel
// ==================================================

// --------------------------
// Payload → Model
// --------------------------

// DriverCancelFromPayload
// ใช้ใน Consumer
// - payload (Kafka)
// - proto.Unmarshal → pb
// - pb → model
func DriverCancelFromPayload(
	payload []byte,
) (model.DriverCancelResponse, error) {
	var pbResp pb.DriverCancelResponse
	if err := proto.Unmarshal(payload, &pbResp); err != nil {
		return model.DriverCancelResponse{}, err
	}

	return DriverCancelFromPB(&pbResp), nil
}

// --------------------------
// PB → Model
// --------------------------

// DriverCancelFromPB
// แปลง protobuf → domain model
func DriverCancelFromPB(
	pbResp *pb.DriverCancelResponse,
) model.DriverCancelResponse {
	return model.DriverCancelResponse{
		OrderID:  pbResp.OrderId,
		DriverID: pbResp.DriverId,
		//Status:    pbResp.Status,
		Timestamp: time.Unix(pbResp.Timestamp, 0),
	}
}

// --------------------------
// Model → PB
// --------------------------

// DriverCancelToPB
// ใช้กรณีที่ระบบนี้ต้องส่ง DriverCancel event ออกไป (future-proof)
func DriverCancelToPB(
	m model.DriverCancelResponse,
) *pb.DriverCancelResponse {
	ts := m.Timestamp
	if ts.IsZero() {
		ts = time.Now()
	}

	return &pb.DriverCancelResponse{
		OrderId:  m.OrderID,
		DriverId: m.DriverID,
		//Status:    m.Status,
		Timestamp: ts.Unix(),
	}
}
