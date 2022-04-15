package model

import (
	validator "gopkg.in/go-playground/validator.v9"
)

// 基础用户信息表
type AdminUser struct {
	AdminId    uint   `json:"admin_id"`    // 用户id
	AdminName  string `json:"admin_name"`  // 管理员用户名
	RealName   string `json:"real_name"`   // 管理员真实姓名
	Password   string `json:"-"`           // 密码(需要加密保存)
	Tel        string `json:"tel"`         // 手机号
	Email      string `json:"email"`       // 邮件
	Permission int    `json:"permission"`  // 权限
	LastTime   string `json:"last_time"`   // 用户的最后一次上线时间
	CreateTime string `json:"create_time"` // 该账户的注册时间
	Birthday   string `json:"birthday"`    // 该账户的生日，用来确定年龄
	Sex        string `json:"sex"`         // 用户性别，可选项"男"，"女"，"其他"
	Status     string `json:"status"`      // 用户状态，可选项"开启"，"关闭"
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
