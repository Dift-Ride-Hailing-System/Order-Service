package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"dift_backend_go/order-service/config"

	// infra
	"dift_backend_go/order-service/internal/integration/cache"
	"dift_backend_go/order-service/internal/integration/kafka"

	// core
	"dift_backend_go/order-service/internal/service/core"

	// flows
	"dift_backend_go/order-service/internal/service/flow_cancel"

	// adapters
	"dift_backend_go/order-service/internal/adapter"

	// http
	route "dift_backend_go/order-service/route"
)

func main() {

	// Load config

	cfg := config.LoadConfig("config/config.yaml")

	// Redis

	redisCache := cache.NewRedisCache(
		cfg.Redis.Addr,
		cfg.Redis.Password,
		cfg.Redis.DB,
	)

	// Kafka Producer

	producer := kafka.NewKafkaWriter(cfg.Kafka.BootstrapServers)
	defer producer.Close()

	// Core Order Service

	orderSvc := core.NewOrderService(
		cfg,
		producer,
		redisCache,
	)

	// Flow Services

	cancelFlow := flow_cancel.New(orderSvc)

	// Adapters

	travelAdapter := adapter.NewTravelConsumer(orderSvc)
	matchAdapter := adapter.NewMatchResultConsumer(orderSvc)
	driverCancelAdapter := adapter.NewDriverCancelConsumer(cancelFlow)

	// Kafka Consumers

	travelConsumer := kafka.NewKafkaReader(
		cfg.Kafka.BootstrapServers,
		cfg.Kafka.GroupID,
		cfg.Kafka.TravelTopic,
	)
	defer travelConsumer.Close()

	matchConsumer := kafka.NewKafkaReader(
		cfg.Kafka.BootstrapServers,
		cfg.Kafka.GroupID,
		cfg.Kafka.ConsumerTopic,
	)
	defer matchConsumer.Close()

	driverCancelConsumer := kafka.NewKafkaReader(
		cfg.Kafka.BootstrapServers,
		cfg.Kafka.GroupID,
		cfg.Kafka.DriverCancelTopic,
	)
	defer driverCancelConsumer.Close()

	// Consume Kafka

	go travelConsumer.Consume(func(_ string, val []byte) error {
		return travelAdapter.Handle(val)
	})

	go matchConsumer.Consume(func(_ string, val []byte) error {
		return matchAdapter.Handle(val)
	})

	go driverCancelConsumer.Consume(func(_ string, val []byte) error {
		return driverCancelAdapter.Handle(val)
	})

	// =============================
	// HTTP Server
	// =============================
	r := gin.Default()
	route.RegisterRoutes(r, orderSvc)

	go func() {
		log.Printf("[Main] API listening on :%s\n", cfg.Server.Port)
		if err := r.Run(":" + cfg.Server.Port); err != nil {
			log.Fatal(err)
		}
	}()

	// =============================
	// Graceful Shutdown
	// =============================
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("[Main] shutting down order-service")
	time.Sleep(1 * time.Second)
}
