package adminctrl

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"github.com/kalougata/go-take-out/internal/model"
	"github.com/kalougata/go-take-out/pkg/errors"
	"github.com/kalougata/go-take-out/pkg/response"
)

type categoryController struct {
}

type CategoryController interface {
	AddCategory(c *fiber.Ctx) error
}

// AddCategory 添加一个分类
func (cc *categoryController) AddCategory(c *fiber.Ctx) error {
	body := &model.CreateCategoryRequest{}
	if err := c.BodyParser(body); err != nil {
		return response.Build(c, errors.ErrBadRequest().WithError(err), err.Error())
	}
	if v := validate.Struct(body); !v.Validate() {
		return response.Build(c, errors.ErrInvalidRequestParams(), v.Errors)
	}

	return response.Build(c, nil, "创建分类")
}

func NewCategoryController() CategoryController {
	return &categoryController{}
}
