package auth

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"gitlab.com/revolutionize1/foward-api/internal/api/database/model"
	"gitlab.com/revolutionize1/foward-api/internal/api/database/postgres"
	"gitlab.com/revolutionize1/foward-api/internal/api/database/redis"
	"gitlab.com/revolutionize1/foward-api/internal/api/middleware/auth/limiter"
	"gitlab.com/revolutionize1/foward-api/internal/api/response"
)

func ValidateKey(c *fiber.Ctx, key string) (bool, error) {
	// Check in Redis first
	apiKey, err := redis.RetrieveApiKey(key)
	if err != nil {
		log.Error(err)
		return false, err
	}

	if apiKey != nil && !verifyApiKeyAddressIp(apiKey, c.IP()) {
		return false, errors.New("invalid ip address provided")
	}

	if apiKey != nil {
		return true, nil
	}

	// If not found in Redis, check in PostgreSQL
	apiKey, err = postgres.RetrieveApiKey(key)
	if err != nil {
		return false, err
	}

	go func(apiKey *model.ApiKey) {
		// If found in PostgreSQL, store in Redis for future checks
		if apiKey != nil {
			err := redis.StoreApiKey(apiKey)
			if err != nil {
				log.Error(err)
			}
		}
	}(apiKey)

	if apiKey != nil && !verifyApiKeyAddressIp(apiKey, c.IP()) {
		return false, errors.New("invalid ip address provided")
	}

	return apiKey != nil, nil
}

func verifyApiKeyAddressIp(apiKey *model.ApiKey, ip string) bool {
	return apiKey.IpAddress == ip
}

func mapErrorToStatusAndMessage(err error) (string, interface{}) {
	switch err {
	case keyauth.ErrMissingOrMalformedAPIKey:
		return "fail", err.Error()
	default:
		return "error", err.Error()
	}
}

func HandleApiKeyError(c *fiber.Ctx, err error) error {
	if err := limiter.Limit(c); err != nil {
		errResponse := response.Response{
			Status:  "fail",
			Message: err.Message,
		}

		return c.Status(err.Code).JSON(&errResponse)
	}

	status, message := mapErrorToStatusAndMessage(err)

	errResponse := response.Response{
		Status:  status,
		Message: message,
	}

	return c.Status(fiber.StatusUnauthorized).JSON(&errResponse)
}
