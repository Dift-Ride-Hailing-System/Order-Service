package core

import (
	"time"

	"dift_backend_go/order-service/config"
	//cache "dift_backend_go/order-service/internal/integration/cache" แก้ไข ลบก่อน กูทำผิด
	"dift_backend_go/order-service/internal/integration/kafka"
	port "dift_backend_go/order-service/internal/interface"
	"dift_backend_go/order-service/internal/service/coupon"
)

type OrderService struct {
	cfg *config.Config

	// infrastructure
	producer kafka.KafkaProducer
	cache    port.Cache
	idem     *Idempotency

	// domain services
	userCouponFetcher *coupon.UserCouponFetcher
	couponPrecalc     *coupon.CouponPrecalculator
	couponApplier     *coupon.CouponApplier

	// background config
	cacheTTL   time.Duration
	cleanupInt time.Duration
}

func NewOrderService(
	cfg *config.Config,
	couponFetcher *coupon.UserCouponFetcher,
	precalc *coupon.CouponPrecalculator,
	applier *coupon.CouponApplier,
	producer kafka.KafkaProducer,
	c port.Cache,
) *OrderService {

	s := &OrderService{
		cfg:               cfg,
		producer:          producer,
		cache:             c,
		idem:              NewIdempotency(cfg.Idempotency.TTL),
		cacheTTL:          cfg.Redis.TTL,
		cleanupInt:        1 * time.Minute,
		userCouponFetcher: couponFetcher,
		couponPrecalc:     precalc,
		couponApplier:     applier,
	}

	// background cleanup owned by core
	go s.cleanupExpiredOrders()

	return s
}

/* ---------- getters (สำคัญมาก) ---------- */

// cache
func (s *OrderService) Cache() port.Cache {
	return s.cache
}

func (s *OrderService) CacheTTL() time.Duration {
	return s.cacheTTL
}

// idempotency
func (s *OrderService) Idem() *Idempotency {
	return s.idem
}

// kafka
func (s *OrderService) Producer() kafka.KafkaProducer {
	return s.producer
}

// coupon services
func (s *OrderService) CouponFetcher() *coupon.UserCouponFetcher {
	return s.userCouponFetcher
}

func (s *OrderService) CouponPrecalc() *coupon.CouponPrecalculator {
	return s.couponPrecalc
}

func (s *OrderService) CouponApplier() *coupon.CouponApplier {
	return s.couponApplier
}
