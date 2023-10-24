package adminctrl

import (
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

func (ac *authController) Login(c *fiber.Ctx) error {
	var data = new(model.EmployeeLoginRequest)
	if err := c.BodyParser(data); err != nil {
		return response.Build(c, errors.ErrInvalidRequestParams().WithMsg(err.Error()), nil)
	}

	v := validate.Struct(data)

	if !v.Validate() {
		return response.Build(c, errors.ErrInvalidRequestParams(), v.Errors)
	}

	if resp, err := ac.service.Login(c.Context(), data); err == nil {
		return response.Build(c, nil, resp)
	} else {
		return response.Build(c, err, nil)
	}
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
