package router

import (
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	App *fiber.App
}

func New() *Router {
	return &Router{
		App: fiber.New(),
	}
}
