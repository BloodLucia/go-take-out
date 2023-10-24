package adminctrl

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"github.com/kalougata/go-take-out/internal/model"
	adminsrv "github.com/kalougata/go-take-out/internal/service/admin"
	"github.com/kalougata/go-take-out/pkg/errors"
	"github.com/kalougata/go-take-out/pkg/response"
)

type authController struct {
	service adminsrv.EmployeeService
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
		return response.Build(c, errors.ErrInvalidRequestParams(), err.Error())
	}

	v := validate.Struct(data)

	if !v.Validate() {
		return response.Build(c, errors.ErrInvalidRequestParams().WithError(v.Errors), v.Errors)
	}

	data.RegIp = c.IP()
	if err := ac.service.Register(c.Context(), data); err != nil {
		return response.Build(c, err, nil)
	}

	return response.Build(c, nil, nil)
}

func NewAuthController(service adminsrv.EmployeeService) AuthController {
	return &authController{service}
}
