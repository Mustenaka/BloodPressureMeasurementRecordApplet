package model

import (
	validator "gopkg.in/go-playground/validator.v9"
)

// BNP检验指标
type TestIndicatorBnps struct {
	Id       uint   `json:"id"`
	UserId   uint   `json:"user_id"`
	Data     int    `json:"data"`      // 检测数据
	CreateAt string `json:"create_at"` // 报告时间
}

// 获取表名称
func (TestIndicatorBnps) TableName() string {
	return "test_indicator_bnps"
}

// 判断有效性
func (testIndicatorBnps *TestIndicatorBnps) Validate() error {
	validate := validator.New()
	return validate.Struct(testIndicatorBnps)
}
