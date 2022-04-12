package model

import (
	validator "gopkg.in/go-playground/validator.v9"
)

// 基础用户信息表
type AdminUser struct {
	AdminId    uint   // 用户id
	AdminName  string // 管理员用户名
	RealName   string // 管理员真实姓名
	Password   string // 密码(需要加密保存)
	Tel        string // 手机号
	Email      string // 邮件
	Permission int    // 权限
	LastTime   string // 用户的最后一次上线时间
	CreateTime string // 该账户的注册时间
	Birthday   string // 该账户的生日，用来确定年龄
	Sex        string // 用户性别，可选项"男"，"女"，"其他"
	Status     string // 用户状态，可选项"开启"，"关闭"
}

// 获取表名称
func (AdminUser) TableName() string {
	return "admin_users"
}

// 判断有效性
func (adminUser *AdminUser) Validate() error {
	validate := validator.New()
	return validate.Struct(adminUser)
}
