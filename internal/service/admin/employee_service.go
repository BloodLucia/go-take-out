package adminsrv

import (
	"context"
	"errors"
	"time"

	"github.com/kalougata/go-take-out/internal/model"
	"github.com/kalougata/go-take-out/pkg/utils"
	"gorm.io/gorm"
)

type employeeService struct {
	*Service
}

type EmployeeService interface {
	Register(ctx context.Context, req *model.EmployeeRegisterRequest) error
	LoginByEmailOrLoginName(ctx context.Context) error
}

// Register implements EmployeeService.
func (es *employeeService) Register(ctx context.Context, req *model.EmployeeRegisterRequest) error {
	employee := &model.Employee{}
	if err := es.DB.WithContext(ctx).
		Model(employee).
		Where("login_name = ?", req.LoginName).
		First(employee).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("登录名已被注册, 请重新输入")
	}
	employee.LoginName = req.LoginName
	employee.Passwd = utils.BcryptHash(req.Passwd)
	employee.RegTime = time.Now()
	employee.RegIp = req.RegIp

	return es.DB.WithContext(ctx).Create(employee).Error
}

// LoginByEmailOrLoginName implements EmployeeService.
func (es *employeeService) LoginByEmailOrLoginName(ctx context.Context) error {
	return nil
}

func NewEmployeeService(service *Service) EmployeeService {
	return &employeeService{Service: service}
}
