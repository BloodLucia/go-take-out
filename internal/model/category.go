package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name       string `gorm:"type:varchar(32);not null;column:name;comment:分类名称"`
	Type       int    `gorm:"type:tinyint(1);null;default:null;column:type;comment:分类类型(0:菜品分类 1:套餐分类)"`
	Sort       int    `gorm:"type:tinyint(1);not null;default:0;column:sort;comment:分类排序(数字越小越靠前)"`
	Status     int    `gorm:"type:tinyint(1);not null;default:1;column:status;comment:分类状态(0:禁用 1:正常)"`
	CreateUser uint   `gorm:"type:bigint;not null;column:create_user;comment:创建人ID"`
	UpdateUser uint   `gorm:"type:bigint;not null;column:update_user;comment:更新人ID"`
}

type CreateCategoryRequest struct {
	Name       string `json:"name" validate:"required|max_len:8" message:"required:name 必填|max_len:name 最多不能超过8个字符"`
	Sort       int    `json:"sort" validate:"required|int" message:"required:sort 必填|int:sort 必须是数字类型"`
	CreateUser string `json:"create_user" validate:"required" message:"required:create_user 必填"`
}

func (c *Category) TableName() string {
	return "tb_categories"
}
