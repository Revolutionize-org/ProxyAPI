package limiter

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/revolutionize1/foward-api/internal/api/database/redis"
)

func Limit(c *fiber.Ctx, config ...*config) *fiber.Error {
	cfg := getConfig(config...)

	ip := c.IP()
	err := redis.StoreFailedAttempt(ip)

	if err != nil {
		// Redis error
		return fiber.ErrInternalServerError
	}

	if isAllowed(cfg, ip) {
		return nil
	}

	return fiber.ErrTooManyRequests
}

func isAllowed(cfg *config, ip string) bool {
	attempts, err := redis.RetrieveFailedAttempt(ip)

	if err != nil {
		return false
	}

	if attempts >= cfg.Max {
		return false
	}
	return true
}
