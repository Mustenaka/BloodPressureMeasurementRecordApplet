package mysql

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"BloodPressure/pkg/db"
	timeconvert "BloodPressure/tools/timeConvert"
	"context"
	"time"
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

// 获取指定日期（截止到当日）记录组
func (ur *patientBpRecordRepo) GetRecordByIdLimitDays(ctx context.Context, id uint, limitdays int) ([]model.PatientBpRecord, error) {
	// 获取目标时间段
	period := time.Now().AddDate(0, 0, -limitdays)
	periodStr := timeconvert.DateString(period)
	records := []model.PatientBpRecord{}
	err := ur.ds.Master().Where("user_id = ? and record_date >= ?", id, periodStr).Find(&records).Error
	return records, err
}

// 添加一个血压记录
func (ur *patientBpRecordRepo) AddRecord(ctx context.Context, id uint, low, high, heartRate int) error {
	nowDate := timeconvert.NowDateString()
	nowTime := timeconvert.NowTimeString()
	records := &model.PatientBpRecord{
		UserId:       id,
		RecordDate:   nowDate,
		RecordTime:   nowTime,
		LowPressure:  low,
		HighPressure: high,
		HeartRate:    heartRate,
	}
	err := ur.ds.Master().Create(records).Error
	return err
}

// 添加一个血压记录
func (ur *patientBpRecordRepo) AddRecordWithDateTime(ctx context.Context, date, time string, id uint, low, high, heartRate int) error {
	records := &model.PatientBpRecord{
		UserId:       id,
		RecordDate:   date,
		RecordTime:   time,
		LowPressure:  low,
		HighPressure: high,
		HeartRate:    heartRate,
	}
	err := ur.ds.Master().Create(records).Error
	return err
}
