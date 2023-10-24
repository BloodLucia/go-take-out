package adminctrl

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"github.com/kalougata/go-take-out/internal/data"
)

type authController struct {
	data *data.Data
}

type AuthController interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type AdminLoginRequest struct {
	LoginName string `json:"login_name" validate:"required" message:"required:login_name 必填"`
	Passwd    string `json:"passwd" validate:"required" message:"required:passwd 必填"`
}

func (ac *authController) Login(c *fiber.Ctx) error {
	var req = new(AdminLoginRequest)
	if err := c.BodyParser(req); err != nil {
		c.SendStatus(http.StatusUnprocessableEntity)
		return c.JSON(fiber.Map{
			"errs": err.Error(),
		})
	}

	v := validate.Struct(req)
	if !v.Validate() {
		c.SendStatus(http.StatusUnprocessableEntity)
		return c.JSON(fiber.Map{
			"errs": v.Errors,
		})
	}

	return c.SendString("Login")
}

func (ac *authController) Register(c *fiber.Ctx) error {
	return c.SendString("register")
}

func NewAuthController(data *data.Data) AuthController {
	return &authController{data}
}
