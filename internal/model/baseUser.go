package model

import (
	validator "gopkg.in/go-playground/validator.v9"
)

// 基础用户信息表
type BaseUser struct {
	UserId     uint   // 用户id
	OpenId     string // 用户的微信openid
	UserName   string // 用户名称-真实姓名
	RealName   string // 用户真实姓名
	Password   string // 用户密码(需要加密保存)
	Tel        string // 用户手机号
	Email      string // 用户邮件
	Permission int    // 用户权限
	LastTime   string // 用户的最后一次上线时间
	CreateTime string // 该账户的注册时间
	Birthday   string // 该账户的生日，用来确定年龄
	Sex        string // 用户性别，可选项"男"，"女"，"其他"
	Status     string // 用户状态，可选项"开启"，"关闭"
}

// 获取表名称
func (BaseUser) TableName() string {
	return "base_users"
}

// 判断有效性
func (baseUser *BaseUser) Validate() error {
	validate := validator.New()
	return validate.Struct(baseUser)
}
