package model

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	LoginName string `gorm:"type:varchar(50) column:login_name comment:用户登录账号"`
	Email     string `gorm:"type:varchar(100) column:email comment:用户邮箱"`
	Passwd    string `gorm:"type:varchar(60) column:passwd comment:加密后的密码"`
	Status    int    `gorm:"type tinyint(1) default 0 comment:用户状态(0-正常 1-锁定)"`
}

type EmployeeRegisterRequest struct {
	Email     string `json:"email" validate:"required|email" message:"required:email 邮箱必填|email:邮箱格式错误"`
	LoginName string `json:"login_name" validate:"required|min_len:5|max_len:10" message:"required:login_name 必填"`
	Passwd    string `json:"passwd" validate:"required|min_len:6|max_len:20" message:"required:passwd 必填"`
}

type EmployeeLoginRequest struct {
	LoginName string `json:"login_name" validate:"requried" message:"required:login_name 必填"`
	Passwd    string `json:"passwd" validate:"required" message:"required:passwd 必填"`
}

func (Employee) TableName() string {
	return "tb_employees"
}
