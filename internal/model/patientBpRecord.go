package model

import validator "gopkg.in/go-playground/validator.v9"

// 血压记录表
type PatientBpRecord struct {
	RecordId     uint   // 血压记录id
	UserId       uint   // 用户id
	RecordDate   string // 血压记录日期
	RecordTime   string // 血压记录时间
	LowPressure  int    // 低压
	HighPressure int    // 高压
	HeartRate    int    // 心率
}

// 获取表名称
func (PatientBpRecord) TableName() string {
	return "patient_bp_records"
}

// 判断有效性
func (patientBpRecord *PatientBpRecord) Validate() error {
	validate := validator.New()
	return validate.Struct(patientBpRecord)
}
