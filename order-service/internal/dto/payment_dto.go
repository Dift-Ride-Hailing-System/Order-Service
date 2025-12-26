package dto

import (
	"time"

	"dift_backend_go/order-service/internal/model"
	pb "dift_backend_go/order-service/proto/pb"
)

// CreatePaymentRequest
func CreatePaymentRequestFromPB(pbMsg *pb.CreatePaymentRequest) model.CreatePaymentRequest {
	return model.CreatePaymentRequest{
		OrderID: pbMsg.OrderId,
		UserID:  pbMsg.UserId,
		Amount:  pbMsg.Amount,
		Method:  model.PaymentMethod(pbMsg.Method),
	}
}

func CreatePaymentRequestToPB(m model.CreatePaymentRequest) *pb.CreatePaymentRequest {
	return &pb.CreatePaymentRequest{
		OrderId: m.OrderID,
		UserId:  m.UserID,
		Amount:  m.Amount,
		Method:  pb.PaymentMethod(m.Method),
	}
}

// CreatePaymentResponse
func CreatePaymentResponseFromPB(pbMsg *pb.CreatePaymentResponse) model.CreatePaymentResponse {
	return model.CreatePaymentResponse{
		PaymentID:   pbMsg.PaymentId,
		Status:      pbMsg.Status,
		QRCodeURL:   pbMsg.QrCodeUrl,
		RedirectURL: pbMsg.RedirectUrl,
		CreatedAt:   time.Unix(pbMsg.CreatedAt, 0),
	}
}

func CreatePaymentResponseToPB(m model.CreatePaymentResponse) *pb.CreatePaymentResponse {
	return &pb.CreatePaymentResponse{
		PaymentId:   m.PaymentID,
		Status:      m.Status,
		QrCodeUrl:   m.QRCodeURL,
		RedirectUrl: m.RedirectURL,
		CreatedAt:   m.CreatedAt.Unix(),
	}
}

// CheckPaymentStatusRequest
func CheckPaymentStatusRequestFromPB(pbMsg *pb.CheckPaymentStatusRequest) model.CheckPaymentStatusRequest {
	return model.CheckPaymentStatusRequest{
		PaymentID: pbMsg.PaymentId,
	}
}

func CheckPaymentStatusRequestToPB(m model.CheckPaymentStatusRequest) *pb.CheckPaymentStatusRequest {
	return &pb.CheckPaymentStatusRequest{
		PaymentId: m.PaymentID,
	}
}

// CheckPaymentStatusResponse
func CheckPaymentStatusResponseFromPB(pbMsg *pb.CheckPaymentStatusResponse) model.CheckPaymentStatusResponse {
	return model.CheckPaymentStatusResponse{
		PaymentID: pbMsg.PaymentId,
		Status:    pbMsg.Status,
		UpdatedAt: time.Unix(pbMsg.UpdatedAt, 0),
	}
}

func CheckPaymentStatusResponseToPB(m model.CheckPaymentStatusResponse) *pb.CheckPaymentStatusResponse {
	return &pb.CheckPaymentStatusResponse{
		PaymentId: m.PaymentID,
		Status:    m.Status,
		UpdatedAt: m.UpdatedAt.Unix(),
	}
}
