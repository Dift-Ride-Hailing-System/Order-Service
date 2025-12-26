package api

import (
	"net/http"

	service "dift_backend_go/order-service/internal/service/core"
	"dift_backend_go/order-service/internal/service/flow_confirm"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, svc *service.OrderService) {

	group := r.Group("/api/v1/orders")

	// --------------------------------------
	// POST /confirm
	// ผู้โดยสารกด "ยืนยันการเดินทาง" หลังเลือกคูปอง
	// --------------------------------------
	group.POST("/confirm", func(c *gin.Context) {

		var req struct {
			OrderID       string `json:"order_id"`
			CouponCode    string `json:"coupon_code"`
			PaymentMethod string `json:"payment_method"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
			return
		}

		// Basic validate
		if req.OrderID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "order_id is required"})
			return
		}
		if req.PaymentMethod == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "payment_method is required"})
			return
		}

		// Call service
		//result, err := flow_confirm.ConfirmFlow(
			//c.Request.Context(),
			//req.OrderID,
			//req.CouponCode,
			//req.PaymentMethod,
		//)

		//if err != nil {
			//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			//return
		//}

		// Response ไป UI
		//c.JSON(http.StatusOK, gin.H{
			//"order_id":       result.OrderID,
			//"pickup":         result.Pickup,
			//"dropoff":        result.Dropoff,
			//"estimated_fare": result.Estimated,
			//"final_total":    result.FinalTotal,
			//"discount":       result.Discount,
			//"coupon_code":    result.CouponCode,
			//"payment_method": result.PaymentMethod,
			//"status":         result.Status, // matching_pending
		//})
	//})
//}
