package adminv1

import (
	"github.com/google/wire"
	adminctrl "github.com/kalougata/go-take-out/internal/controller/admin"
)

type AdminAPIRouter struct {
	authCtrl     adminctrl.AuthController
	categoryCtrl adminctrl.CategoryController
}

func NewAdminAPIRouter(
	authCtrl adminctrl.AuthController,
	categoryCtrl adminctrl.CategoryController,
) *AdminAPIRouter {
	return &AdminAPIRouter{authCtrl, categoryCtrl}
}

var AdminAPIRouterProvider = wire.NewSet(
	adminctrl.NewAuthController,
	adminctrl.NewCategoryController,
	NewAdminAPIRouter,
)
