package app

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/revolutionize1/foward-api/internal/config"
)

var Instance *Application

type Application struct {
	Config   *config.Config
	FiberApp *fiber.App
	Version  string
}

func New() (*Application, error) {
	cfg, err := config.New()

	if err != nil {
		return nil, err
	}

	return &Application{
		Config:   cfg,
		FiberApp: fiber.New(),
		Version:  "v1",
	}, nil
}
