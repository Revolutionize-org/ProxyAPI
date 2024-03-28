package router

import (
	"fmt"

	"gitlab.com/revolutionize1/foward-api/internal/api/service/request"
	"gitlab.com/revolutionize1/foward-api/internal/app"
)

func (r *Router) SetupRequestRoutes() *Router {
	r.App.Post("/send-request", request.ReceiveRequest)
	return r
}

func (r *Router) MountRoute(application *app.Application) *Router {
	mountUri := fmt.Sprintf("/%s/", application.Version)

	application.FiberApp.Mount(mountUri, r.App)
	return r
}
