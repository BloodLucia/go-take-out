package server

import (
	"github.com/gofiber/fiber/v2"
	adminv1 "github.com/kalougata/go-take-out/api/admin"
	"github.com/kalougata/go-take-out/internal/middleware"
)

func NewHTTPServer(
	aar *adminv1.AdminAPIRouter,
	jm *middleware.JWTMiddleware,
) *fiber.App {
	app := fiber.New()

	// adminGroup 后台路由组
	adminGroup := app.Group("/api/v1/admin")
	{
		// 不需要登录的路由
		noAuthGroup := adminGroup.Group("")
		aar.SetupGuestAPIRouter(noAuthGroup)

		// 需要登录的路由
		needAuthGroup := adminGroup.Group("")
		// needAuthGroup.Use(jm.JWTAdmin())
		aar.SetupCategoryAPIRouter(needAuthGroup)
	}

	return app
}
