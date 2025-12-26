package model

import "errors"

// OrderStatus type (optional แต่แนะนำ)
type OrderStatus string

const (
	StatusWaitingCoupon  OrderStatus = "waiting_coupon"
	StatusMatchingSent   OrderStatus = "matching_sent"
	StatusMatched        OrderStatus = "matched"
	StatusDriverAssigned OrderStatus = "driver_assigned"
	StatusCompleted      OrderStatus = "completed"
	StatusCancelled      OrderStatus = "cancelled"
	StatusNotMatched     OrderStatus = "not_matched"
	StatusTimeout        OrderStatus = "timeout"
)

// state machine ของ Order
var validTransitions = map[OrderStatus][]OrderStatus{
	StatusWaitingCoupon:  {StatusMatchingSent, StatusCancelled},
	StatusMatchingSent:   {StatusMatched, StatusDriverAssigned, StatusCancelled, StatusNotMatched, StatusTimeout},
	StatusMatched:        {StatusDriverAssigned, StatusCompleted, StatusCancelled},
	StatusDriverAssigned: {StatusCompleted, StatusCancelled},
	StatusCompleted:      {},
	StatusCancelled:      {},
	StatusNotMatched:     {},
	StatusTimeout:        {},
}

// ValidateOrderStatusTransition ตรวจสอบการเปลี่ยนสถานะ
func ValidateOrderStatusTransition(current, next OrderStatus) error {
	if current == next {
		return nil
	}

	nextList, ok := validTransitions[current]
	if !ok {
		return errors.New("unknown current status: " + string(current))
	}

	for _, v := range nextList {
		if v == next {
			return nil
		}
	}

	return errors.New(
		"invalid status transition from " + string(current) + " to " + string(next),
	)
}
