package adminv1

import "github.com/gofiber/fiber/v2"

func (api *AdminAPIRouter) SetupCategoryAPIRouter(r fiber.Router) {
	categoryGroup := r.Group("/category")
	{
		categoryGroup.Get("list", func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"data": nil,
			})
		})
		categoryGroup.Post("/add", api.categoryCtrl.AddCategory)
	}

}
