package model

import (
	validator "gopkg.in/go-playground/validator.v9"
)

// 基础用户信息表
type BaseUser struct {
	UserId     uint   `json:"user_id"`     // 用户id
	OpenId     string `json:"open_id"`     // 用户的微信openid
	UserName   string `json:"user_name"`   // 用户名称-真实姓名
	RealName   string `json:"real_name"`   // 用户真实姓名
	Password   string `json:"password"`    // 用户密码(需要加密保存)
	Tel        string `json:"tel"`         // 用户手机号
	Email      string `json:"email"`       // 用户邮件
	LastTime   string `json:"last_time"`   // 用户的最后一次上线时间
	CreateTime string `json:"create_time"` // 该账户的注册时间
	Birthday   string `json:"birthday"`    // 该账户的生日，用来确定年龄
	Sex        string `json:"sex"`         // 用户性别，可选项"男"，"女"，"其他"
	Status     string `json:"status"`      // 用户状态，可选项"开启"，"关闭"
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
