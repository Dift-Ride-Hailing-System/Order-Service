package core

import (
	"context"
	"encoding/json"
	"time"

	"dift_backend_go/order-service/internal/model"
	"dift_backend_go/order-service/internal/observability/logger"
)

// cleanupExpiredOrders ลบ order ที่ไม่มีการเคลื่อนไหวเกิน cacheTTL
func (s *OrderService) cleanupExpiredOrders() {
	ticker := time.NewTicker(s.cleanupInt)
	defer ticker.Stop()

	for range ticker.C {
		ctx := context.Background()

		keys, err := s.cache.Keys(ctx, "order:*")
		if err != nil {
			logger.Error("cleanupExpiredOrders: cache.Keys failed: %v", err)
			continue
		}

		for _, key := range keys {
			data, err := s.cache.Get(ctx, key)
			if err != nil || data == nil {
				continue
			}

			var order model.Order
			if err := json.Unmarshal(data, &order); err != nil {
				logger.Error("cleanupExpiredOrders: json.Unmarshal failed for key %s: %v", key, err)
				continue
			}

			// ถ้าไม่มีการอัพเดทเกิน cacheTTL - ลบออก
			if time.Since(order.UpdatedAt) > s.cacheTTL {
				_ = s.cache.Delete(ctx, key)
				logger.Info("cleanupExpiredOrders: expired order removed: %s", order.OrderID)
			}
		}
	}
}
