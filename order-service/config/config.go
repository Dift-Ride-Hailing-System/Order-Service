package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

//
// =====================
//   CONFIG STRUCTS
// =====================
//

// ---------------------
// Server
// ---------------------
type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

// ---------------------
// Kafka
// ---------------------
type KafkaConfig struct {
	BootstrapServers []string `mapstructure:"bootstrap_servers"`
	GroupID          string   `mapstructure:"group_id"`
	Timeout          time.Duration

	// consume topics
	TravelRequestTopic string `mapstructure:"travel_request_topic"`
	MatchResultTopic   string `mapstructure:"match_result_topic"`
	DriverCancelTopic  string `mapstructure:"driver_cancel_topic"`

	// produce topics
	OrderMatchingTopic   string `mapstructure:"order_matching_topic"`
	PassengerCancelTopic string `mapstructure:"passenger_cancel_topic"`
	TripHistoryTopic     string `mapstructure:"trip_history_topic"`
}

// ---------------------
// Redis
// ---------------------
type RedisConfig struct {
	Addr     string        `mapstructure:"addr"`
	Password string        `mapstructure:"password"`
	DB       int           `mapstructure:"db"`
	TTL      time.Duration `mapstructure:"ttl"`
	Prefix   string        `mapstructure:"prefix"`
}

// ---------------------
// Retry
// ---------------------
type RetryConfig struct {
	MaxAttempts int           `mapstructure:"max_attempts"`
	Interval    time.Duration `mapstructure:"interval"`
}

// ---------------------
// gRPC Services
// ---------------------
type GRPCServiceConfig struct {
	Addr string `mapstructure:"addr"`
}

type GRPCConfig struct {
	Timeout           time.Duration     `mapstructure:"timeout"`
	UserCouponService GRPCServiceConfig `mapstructure:"user_coupon_service"`
	CouponService     GRPCServiceConfig `mapstructure:"coupon_service"`
}

// ---------------------
// Payment (HTTP)
// ---------------------
type PaymentServiceConfig struct {
	BaseURL string        `mapstructure:"base_url"`
	Timeout time.Duration `mapstructure:"timeout"`
	Retry   RetryConfig   `mapstructure:"retry"`
}

// ---------------------
// Idempotency
// ---------------------
type IdempotencyConfig struct {
	TTL time.Duration `mapstructure:"ttl"`
}

// ---------------------
// Root Config
// ---------------------
type Config struct {
	Server      ServerConfig         `mapstructure:"server"`
	Kafka       KafkaConfig          `mapstructure:"kafka"`
	Redis       RedisConfig          `mapstructure:"redis"`
	GRPC        GRPCConfig           `mapstructure:"grpc"`
	Payment     PaymentServiceConfig `mapstructure:"payment_service"`
	Idempotency IdempotencyConfig    `mapstructure:"idempotency"`
	Retry       RetryConfig          `mapstructure:"retry"`
}

//
// =====================
//   LOAD CONFIG
// =====================
//

func LoadConfig(path string) *Config {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("failed to unmarshal config: %v", err)
	}

	parseDuration := func(key string, def time.Duration) time.Duration {
		raw := viper.GetString(key)
		if raw == "" {
			return def
		}
		d, err := time.ParseDuration(raw)
		if err != nil {
			log.Printf("[WARN] invalid duration '%s', using default %v", raw, def)
			return def
		}
		return d
	}

	// ---------------------
	// Parse durations
	// ---------------------

	cfg.Kafka.Timeout = parseDuration("kafka.timeout", 10*time.Second)

	cfg.Redis.TTL = parseDuration("redis.ttl", 5*time.Minute)

	cfg.GRPC.Timeout = parseDuration("grpc.timeout", 5*time.Second)

	cfg.Payment.Timeout = parseDuration("payment_service.timeout", 5*time.Second)

	cfg.Idempotency.TTL = parseDuration("idempotency.ttl", 30*time.Second)

	cfg.Retry.Interval = parseDuration("retry.interval", 300*time.Millisecond)
	cfg.Payment.Retry.Interval = parseDuration("payment_service.retry.interval", 200*time.Millisecond)

	// ---------------------
	// Defaults
	// ---------------------

	if cfg.Server.Port == "" {
		cfg.Server.Port = "8080"
	}

	// Kafka defaults
	if len(cfg.Kafka.BootstrapServers) == 0 {
		cfg.Kafka.BootstrapServers = []string{"localhost:9092"}
	}
	if cfg.Kafka.GroupID == "" {
		cfg.Kafka.GroupID = "order-service-group"
	}

	// Retry defaults
	if cfg.Retry.MaxAttempts == 0 {
		cfg.Retry.MaxAttempts = 3
	}
	if cfg.Payment.Retry.MaxAttempts == 0 {
		cfg.Payment.Retry.MaxAttempts = 3
	}

	return &cfg
}
