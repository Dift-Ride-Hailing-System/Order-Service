package model

import "time"

// PaymentMethod ใช้ใน logic ภายใน
type PaymentMethod int32

const (
	PaymentUnknown    PaymentMethod = 0
	PaymentCreditCard PaymentMethod = 1
	PaymentDebitCard  PaymentMethod = 2
	PaymentPromptPay  PaymentMethod = 3
	PaymentWallet     PaymentMethod = 4
	PaymentCash       PaymentMethod = 5
)

// CreatePaymentRequest model
type CreatePaymentRequest struct {
	OrderID string
	UserID  string
	Amount  float64
	Method  PaymentMethod
}

// CreatePaymentResponse model
type CreatePaymentResponse struct {
	PaymentID   string
	Status      string
	QRCodeURL   string
	RedirectURL string
	CreatedAt   time.Time
}

// CheckPaymentStatusRequest model
type CheckPaymentStatusRequest struct {
	PaymentID string
}

// CheckPaymentStatusResponse model
type CheckPaymentStatusResponse struct {
	PaymentID string
	Status    string
	UpdatedAt time.Time
}
