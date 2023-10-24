package adminv1

import (
	"github.com/google/wire"
	adminctrl "github.com/kalougata/go-take-out/internal/controller/admin"
)

type AdminAPIRouter struct {
	authCtrl adminctrl.AuthController
}

func NewAdminAPIRouter(
	authCtrl adminctrl.AuthController,
) *AdminAPIRouter {
	return &AdminAPIRouter{authCtrl}
}

var AdminAPIRouterProvider = wire.NewSet(
	adminctrl.NewAuthController,
	NewAdminAPIRouter,
)
