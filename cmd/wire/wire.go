//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	adminv1 "github.com/kalougata/go-take-out/api/admin"
	"github.com/kalougata/go-take-out/internal/server"
)

func NewApp() (*fiber.App, func(), error) {
	panic(wire.Build(
		adminv1.AdminAPIRouterProvider,
		server.ServerProvider,
	))
}
