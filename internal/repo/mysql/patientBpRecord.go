package mysql

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"BloodPressure/pkg/db"
	timeconvert "BloodPressure/tools/timeConvert"
	"context"
)

var _ repo.PatientBpRecordRepo = (*patientBpRecordRepo)(nil)

type patientBpRecordRepo struct {
	ds db.IDataSource
}

// 创建一个新的UserRepo
func NewPatientBpRecordRepo(_ds db.IDataSource) *patientBpRecordRepo {
	return &patientBpRecordRepo{
		ds: _ds,
	}
}

// 通过id获取全部记录
func (ur *patientBpRecordRepo) GetRecordById(ctx context.Context, id uint) ([]model.PatientBpRecord, error) {
	records := []model.PatientBpRecord{}
	err := ur.ds.Master().Where("user_id = ?", id).Find(&records).Error
	return records, err
}

// 添加一个血压记录
func (ur *patientBpRecordRepo) AddRecord(ctx context.Context, id uint, low, high int) error {
	nowDate := timeconvert.NowDateString()
	nowTime := timeconvert.NowTimeString()
	records := &model.PatientBpRecord{
		UserId:       id,
		RecordDate:   nowDate,
		RecordTime:   nowTime,
		LowPressure:  low,
		HighPressure: high,
	}
	err := ur.ds.Master().Create(records).Error
	return err
}
