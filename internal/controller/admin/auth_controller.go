package adminctrl

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"github.com/kalougata/go-take-out/internal/data"
	"github.com/kalougata/go-take-out/internal/model"
	"github.com/kalougata/go-take-out/pkg/utils"
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
	var data = new(model.EmployeeRegisterRequest)
	if err := c.BodyParser(data); err != nil {
		c.SendStatus(http.StatusUnprocessableEntity)
		return c.JSON(fiber.Map{
			"success": false,
			"errs":    err.Error(),
		})
	}

	employee := &model.Employee{
		LoginName: data.LoginName,
		Email:     data.Email,
		Passwd:    utils.BcryptHash(data.Passwd),
	}

	if err := ac.data.DB.WithContext(c.Context()).Create(employee).Error; err != nil {
		c.SendStatus(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"success": false,
			"errs":    err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"errs":    nil,
		"data":    data,
	})
}

func NewAuthController(data *data.Data) AuthController {
	return &authController{data}
}
