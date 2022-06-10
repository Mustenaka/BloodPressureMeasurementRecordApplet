package model

import (
	validator "gopkg.in/go-playground/validator.v9"
)

// 肌酐检验指标
type TestIndicatorCreatinines struct {
	Id       uint   `json:"id"`
	UserId   uint   `json:"user_id"`
	Data     int    `json:"data"`      // 检测数据
	CreateAt string `json:"create_at"` // 报告时间
}

// 获取表名称
func (TestIndicatorCreatinines) TableName() string {
	return "test_indicator_creatinines"
}

// 判断有效性
func (testIndicatorCreatinines *TestIndicatorCreatinines) Validate() error {
	validate := validator.New()
	return validate.Struct(testIndicatorCreatinines)
}
