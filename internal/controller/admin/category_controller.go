package adminctrl

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalougata/go-take-out/pkg/response"
)

type categoryController struct {
}

type CategoryController interface {
	AddCategory(c *fiber.Ctx) error
}

// AddCategory 添加一个分类
func (cc *categoryController) AddCategory(c *fiber.Ctx) error {
	return response.Build(c, nil, "添加分类")
}

func NewCategoryController() CategoryController {
	return &categoryController{}
}
