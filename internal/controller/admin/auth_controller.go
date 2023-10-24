package adminctrl

import "github.com/gofiber/fiber/v2"

type AuthController struct {
}

func (ac *AuthController) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Login")
	}
}

func (ac *AuthController) Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Register")
	}
}

func NewAuthController() *AuthController {
	return &AuthController{}
}
