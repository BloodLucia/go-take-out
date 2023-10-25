package adminv1

import "github.com/gofiber/fiber/v2"

func (api *AdminAPIRouter) SetupCategoryAPIRouter(r fiber.Router) {
	r.Get("list", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": nil,
		})
	})
	r.Post("add", api.categoryCtrl.AddCategory)
}
