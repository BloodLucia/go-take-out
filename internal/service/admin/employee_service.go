package adminsrv

import (
	"context"
	"errors"

	"time"

	"github.com/kalougata/go-take-out/internal/model"
	myErr "github.com/kalougata/go-take-out/pkg/errors"
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
		return myErr.ErrBadRequest().WithMsg("登录名已被注册, 请重新输入")
	}
	employee.LoginName = req.LoginName
	employee.Passwd = utils.BcryptHash(req.Passwd)
	employee.RegTime = time.Now()
	employee.RegIp = req.RegIp

	if err := es.DB.WithContext(ctx).Create(employee).Error; err != nil {
		return myErr.ErrInternalServer().WithMsg("注册失败, 请稍后再试").WithError(err)
	}

	return nil
}

// LoginByEmailOrLoginName implements EmployeeService.
func (es *employeeService) LoginByEmailOrLoginName(ctx context.Context) error {
	return nil
}

func NewEmployeeService(service *Service) EmployeeService {
	return &employeeService{Service: service}
}
