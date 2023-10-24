package adminctrl

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
}

type AdminLoginRequest struct {
	LoginName string `json:"login_name"`
	Passwd    string `json:"passwd"`
}

func (ac *AuthController) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req = new(AdminLoginRequest)
		if err := c.BodyParser(req); err != nil {
			c.SendStatus(http.StatusUnprocessableEntity)
			return c.JSON(fiber.Map{
				"errs": err.Error(),
			})
		}
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
