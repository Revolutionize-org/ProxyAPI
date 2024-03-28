package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gitlab.com/revolutionize1/foward-api/internal/api/middleware/auth"
)

func UseLogger(app *fiber.App, file *os.File) {
	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${method} | ${path} | ${ip} | ${status} | ${latency}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		Output:     file,
	}))
}

func UseKeyAuth(app *fiber.App) {
	app.Use(keyauth.New(keyauth.Config{
		Validator:    auth.ValidateKey,
		KeyLookup:    "header:X-API-KEY",
		ErrorHandler: auth.HandleApiKeyError,
	}))
}
