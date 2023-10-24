package model

import (
	"time"

	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	LoginName     string    `gorm:"type:varchar(50);not null;unique_index;column:login_name;comment:用户登录账号"`
	Passwd        string    `gorm:"type:varchar(60);not null;column:passwd;comment:加密后的密码"`
	Status        int       `gorm:"type:tinyint(1);default:1;column:status;comment:用户状态(0-正常 1-锁定)"`
	RegIp         string    `gorm:"type:varchar(15);default:null;column:reg_ip;comment:注册IP"`
	RegTime       time.Time `gorm:"type:datetime;null;default:null;column:reg_time;comment:注册时间"`
	LastLoginIp   string    `gorm:"type:varchar(15);null;default:null;column:last_login_ip;comment:最后一次登录的IP"`
	LastLoginTime time.Time `gorm:"type:datetime;null;default:null;column:last_login_time;comment:用户最后登录时间"`
}

type EmployeeRegisterRequest struct {
	LoginName string `json:"login_name" validate:"required|min_len:5|max_len:10" message:"required:login_name 必填"`
	Passwd    string `json:"passwd" validate:"required|min_len:6|max_len:20" message:"required:passwd 必填"`
	RegIp     string `json:"-"`
}

type EmployeeLoginRequest struct {
	LoginName string `json:"login_name" validate:"required" message:"required:login_name 必填"`
	Passwd    string `json:"passwd" validate:"required" message:"required:passwd 必填"`
}

type EmployeeLoginResponse struct {
	UserId    string `json:"id"`
	LoginName string `json:"login_name"`
	Token     string `json:"token"`
}

func (e *Employee) TableName() string {
	return "tb_employees"
}

func (e *Employee) StringID() string {
	return cast.ToString(e.ID)
}
