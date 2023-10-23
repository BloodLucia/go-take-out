package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

type Server struct {
	HTTPServer *fiber.App
}

func NewServer(httpServer *fiber.App) *Server {
	return &Server{HTTPServer: httpServer}
}

var ServerProvider = wire.NewSet(
	NewHTTPServer,
	NewServer,
)
