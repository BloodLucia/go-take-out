package adminv1

import "github.com/gofiber/fiber/v2"

func (api *AdminAPIRouter) SetupGuestAPIRouter(r fiber.Router) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello admin")
	})
	r.Post("/login", api.authCtrl.Login)
	r.Post("/register", api.authCtrl.Register)
}
