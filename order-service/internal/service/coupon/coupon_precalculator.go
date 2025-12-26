package coupon

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	port "dift_backend_go/order-service/internal/interface"
	gprc_port "dift_backend_go/order-service/internal/interface/gprc"
	"dift_backend_go/order-service/internal/model"
	"dift_backend_go/order-service/internal/observability/logger"
)

const couponPrecomputeTimeout = 3 * time.Second

// CouponPrecalculator เป็น service สำหรับคำนวณราคาของคูปองล่วงหน้า
type CouponPrecalculator struct {
	couponService gprc_port.CouponService
	cache         port.Cache
	cacheTTL      time.Duration
}

// NewCouponPrecalculator constructor
func NewCouponPrecalculator(
	couponService gprc_port.CouponService,
	cache port.Cache,
	ttl time.Duration,
) *CouponPrecalculator {
	return &CouponPrecalculator{
		couponService: couponService,
		cache:         cache,
		cacheTTL:      ttl,
	}
}

// PrecomputePrices - คำนวณราคาของคูปองทั้งหมดและเก็บลง cache
func (p *CouponPrecalculator) PrecomputePrices(
	ctx context.Context,
	userID string,
	orderTotal float64,
	coupons []string,
) error {

	if len(coupons) == 0 {
		return errors.New("no coupons")
	}

	// สร้าง context สำหรับ timeout
	ctx, cancel := context.WithTimeout(ctx, couponPrecomputeTimeout)
	defer cancel()

	results := make(map[string]model.CouponPriceResult)

	for _, code := range coupons {
		input := model.ApplyCouponInput{
			UserID:     userID,
			CouponCode: code,
			OrderTotal: orderTotal,
		}

		// เรียก business logic ผ่าน interface CouponService
		result, err := p.couponService.ApplyCoupon(ctx, input)
		if err != nil {
			logger.Warn(
				"CouponPrecalculator: failed coupon=%s err=%v",
				code,
				err,
			)
			continue
		}

		// เก็บผลลัพธ์ลง map
		results[code] = model.CouponPriceResult{
			FinalTotal: result.FinalTotal,
			Discount:   result.Discount,
		}
	}

	if len(results) == 0 {
		return errors.New("no coupon price computed")
	}

	// แปลง map เป็น JSON เพื่อเก็บลง cache
	data, err := json.Marshal(results)
	if err != nil {
		return err
	}

	key := "precomputed_prices:" + userID

	// เก็บลง cache พร้อม TTL
	return p.cache.Set(ctx, key, data, int(p.cacheTTL.Seconds()))
}
