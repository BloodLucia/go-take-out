package adminsrv

import (
	"github.com/google/wire"
	"github.com/kalougata/go-take-out/internal/data"
)

type Service struct {
	*data.Data
}

func NewService(data *data.Data) *Service {
	return &Service{Data: data}
}

var AdminServiceProvider = wire.NewSet(
	NewService,
	NewEmployeeService,
)
