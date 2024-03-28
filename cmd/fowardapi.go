package main

import (
	"github.com/gofiber/fiber/v2/log"
	"gitlab.com/revolutionize1/foward-api/internal/api/database/postgres"
	"gitlab.com/revolutionize1/foward-api/internal/api/database/redis"
	"gitlab.com/revolutionize1/foward-api/internal/api/middleware"
	"gitlab.com/revolutionize1/foward-api/internal/api/router"
	"gitlab.com/revolutionize1/foward-api/internal/app"
	"gitlab.com/revolutionize1/foward-api/internal/logging"
)

func initializeApp() {
	appInstance, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	app.Instance = appInstance
}

func main() {
	initializeApp()

	if err := logging.Setup(); err != nil {
		log.Fatal(err)
	}

	middleware.UseKeyAuth(app.Instance.FiberApp)

	postgres.Init()
	redis.Init()

	requestRouter := router.New()
	requestRouter.SetupRequestRoutes().MountRoute(app.Instance)

	if err := app.Instance.FiberApp.Listen(app.Instance.Config.Api.Port); err != nil {
		log.Fatal(err)
	}
}
