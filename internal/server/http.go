package server

import (
	"github.com/gofiber/fiber/v2"
	adminv1 "github.com/kalougata/go-take-out/api/admin"
)

func NewHTTPServer(
	aar *adminv1.AdminAPIRouter,
) *fiber.App {
	app := fiber.New()

	// adminGroup 后台路由组
	adminGroup := app.Group("/api/v1/admin")

	// 不需要登录的路由
	noAuthGroup := adminGroup.Group("")
	aar.SetupGuestAPIRouter(noAuthGroup)

	return app
}
