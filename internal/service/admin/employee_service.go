package adminsrv

import (
	"context"

	"github.com/kalougata/go-take-out/internal/data"
	"github.com/kalougata/go-take-out/internal/model"
	"github.com/kalougata/go-take-out/pkg/utils"
)

type employeeService struct {
	data *data.Data
}

type EmployeeService interface {
	Register(ctx context.Context, req *model.EmployeeRegisterRequest) error
	LoginByEmailOrLoginName(ctx context.Context) error
}

// Register implements EmployeeService.
func (es *employeeService) Register(ctx context.Context, req *model.EmployeeRegisterRequest) error {
	employee := &model.Employee{}

	employee.LoginName = req.LoginName
	employee.Email = req.Email
	employee.Passwd = utils.BcryptHash(req.Passwd)

	return es.data.DB.WithContext(ctx).Create(employee).Error
}

// LoginByEmailOrLoginName implements EmployeeService.
func (*employeeService) LoginByEmailOrLoginName(ctx context.Context) error {
	return nil
}

func NewEmployeeService(data *data.Data) EmployeeService {
	return &employeeService{data}
}
