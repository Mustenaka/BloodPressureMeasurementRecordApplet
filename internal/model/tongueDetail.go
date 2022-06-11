package model

import (
	validator "gopkg.in/go-playground/validator.v9"
)

type TongueDetail struct {
	Id            uint   `json:"id"`
	UserId        uint   `json:"user_id"`
	Tongue        string `json:"tongue"`
	TongueCoating string `json:"tongue_coating"`
	Pulse         string `json:"pulse"`
	CreateAt      string `json:"create_at"`
}

// 获取表名称
func (TongueDetail) TableName() string {
	return "tongue_detail"
}

// 判断有效性
func (tongueDetail *TongueDetail) Validate() error {
	validate := validator.New()
	return validate.Struct(tongueDetail)
}
