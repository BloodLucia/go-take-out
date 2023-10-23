//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/kalougata/go-take-out/internal/server"
)

func NewApp() (*fiber.App, func(), error) {
	panic(wire.Build(
		server.ServerProvider,
	))
}
