package response

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	myErrs "github.com/kalougata/go-take-out/pkg/errors"
)

type response struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Data    any    `json:"data"`
	Errs    any    `json:"errs"`
}

func Build(c *fiber.Ctx, err error, data any) error {
	if err == nil {
		c.SendStatus(http.StatusOK)
		return c.JSON(response{
			Code:    http.StatusOK,
			Success: true,
			Msg:     "ok",
			Data:    data,
			Errs:    nil,
		})
	}

	var myErr *myErrs.Error

	if !errors.As(err, &myErr) {
		c.SendStatus(http.StatusInternalServerError)
		return c.JSON(response{
			Code:    http.StatusInternalServerError,
			Success: false,
			Msg:     "Unknown Error",
			Data:    nil,
			Errs:    nil,
		})
	}

	c.SendStatus(myErr.Code)
	return c.JSON(response{
		Code:    myErr.Code,
		Success: false,
		Msg:     myErr.Msg,
		Data:    nil,
		Errs:    data,
	})
}
