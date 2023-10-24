package adminsrv

import (
	"context"
	"errors"

	"time"

	"github.com/kalougata/go-take-out/internal/model"
	myErr "github.com/kalougata/go-take-out/pkg/errors"
	"github.com/kalougata/go-take-out/pkg/jwt"
	"github.com/kalougata/go-take-out/pkg/utils"
	"gorm.io/gorm"
)

type employeeService struct {
	*Service
	jwt *jwt.JWT
}

type EmployeeService interface {
	Register(ctx context.Context, req *model.EmployeeRegisterRequest) error
	Login(ctx context.Context, req *model.EmployeeLoginRequest) (resp *model.EmployeeLoginResponse, err error)
	FindByLoginName(ctx context.Context, loginName string) (employee *model.Employee, exists bool, err error)
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
func (es *employeeService) Login(ctx context.Context, req *model.EmployeeLoginRequest) (resp *model.EmployeeLoginResponse, err error) {
	employee, exists, err := es.FindByLoginName(ctx, req.LoginName)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, myErr.ErrNotFound().WithMsg("登录名不存在, 可能未注册")
	}
	if employee.Status == 1 {
		return nil, myErr.ErrBadRequest().WithMsg("账号已被禁用, 请联系管理员")
	}
	if !utils.BcryptCheck(req.Passwd, employee.Passwd) {
		return nil, myErr.ErrBadRequest().WithMsg("密码错误, 请重新输入")
	}

	employee.LastLoginTime = time.Now()
	if err := es.DB.WithContext(ctx).Model(employee).Save(employee).Error; err != nil {
		return nil, myErr.ErrInternalServer().WithError(err)
	}

	// 生成 token
	claims := jwt.MyCustomClaims{UserId: employee.StringID(), LoginName: employee.LoginName}
	token, err := es.jwt.BuildToken(claims, time.Now().Add(10*time.Minute))
	if err != nil {
		return nil, myErr.ErrInternalServer().WithMsg("生成token失败")
	}
	resp = &model.EmployeeLoginResponse{Token: token}

	return resp, nil
}

// FindByLoginName implements EmployeeService.
func (es *employeeService) FindByLoginName(ctx context.Context, loginName string) (employee *model.Employee, exists bool, err error) {
	employee = &model.Employee{}
	err = es.DB.WithContext(ctx).Model(employee).Where("login_name = ?", loginName).First(employee).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		return nil, false, myErr.ErrInternalServer().WithError(err)
	}

	return employee, true, nil
}

func NewEmployeeService(service *Service, jwt *jwt.JWT) EmployeeService {
	return &employeeService{Service: service, jwt: jwt}
}
