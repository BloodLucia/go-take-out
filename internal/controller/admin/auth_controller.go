package adminctrl

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type authController struct {
}

type AuthController interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type AdminLoginRequest struct {
	LoginName string `json:"login_name"`
	Passwd    string `json:"passwd"`
}

func (ac *authController) Login(c *fiber.Ctx) error {
	var req = new(AdminLoginRequest)
	if err := c.BodyParser(req); err != nil {
		c.SendStatus(http.StatusUnprocessableEntity)
		return c.JSON(fiber.Map{
			"errs": err.Error(),
		})
	}

	return c.SendString("Login")
}

func (ac *authController) Register(c *fiber.Ctx) error {
	return c.SendString("register")
}

func NewAuthController() AuthController {
	return &authController{}
}
