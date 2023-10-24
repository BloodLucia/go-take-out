//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	adminv1 "github.com/kalougata/go-take-out/api/admin"
	"github.com/kalougata/go-take-out/internal/data"
	"github.com/kalougata/go-take-out/internal/server"
	adminsrv "github.com/kalougata/go-take-out/internal/service/admin"
)

func NewApp() (*fiber.App, func(), error) {
	panic(wire.Build(
		data.NewData,
		adminsrv.AdminServiceProvider,
		adminv1.AdminAPIRouterProvider,
		server.ServerProvider,
	))
}
