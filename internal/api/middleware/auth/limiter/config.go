package limiter

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

var instance config

type config struct {
	Max          int64
	Expiration   time.Duration
	KeyGenerator func(*fiber.Ctx) string
}

var defaultConfig = &config{
	Max:        5,
	Expiration: 30 * time.Second,
}

func getConfig(conf ...*config) *config {
	if conf != nil {
		return &instance
	}

	if len(conf) < 1 {
		return defaultConfig
	}

	cfg := conf[0]

	if cfg.Max <= 0 {
		cfg.Max = defaultConfig.Max
	}

	if int(cfg.Expiration.Seconds()) <= 0 {
		cfg.Expiration = defaultConfig.Expiration
	}

	return cfg
}
