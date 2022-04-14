package service

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"context"
)

var _ PatientBpRecordService = (*patientBpRecordService)(nil)

// PatientBpRecordService 定义用户操作服务接口
type PatientBpRecordService interface {
	// 查询
	GetById(ctx context.Context, uid uint) ([]model.PatientBpRecord, error)
	GetByIdLimitDay(ctx context.Context, uid uint, limitdays int) ([]model.PatientBpRecord, error)
	// 添加一条记录
	AddById(ctx context.Context, uid uint, low, high, heartRate int) error
}

// patientBpRecordService 实现UserService接口
type patientBpRecordService struct {
	bprr repo.PatientBpRecordRepo
}

// 新血压记录服务
func NewPBPRecordService(_ur repo.PatientBpRecordRepo) *patientBpRecordService {
	return &patientBpRecordService{
		bprr: _ur,
	}
}

// 获取全部的测量数据
func (us *patientBpRecordService) GetById(ctx context.Context, uid uint) ([]model.PatientBpRecord, error) {
	return us.bprr.GetRecordById(ctx, uid)
}

// 根据限制天数获取数据
func (us *patientBpRecordService) GetByIdLimitDay(ctx context.Context, uid uint, limitdays int) ([]model.PatientBpRecord, error) {
	// 如果限制天数为0，则降级为获取全部数据
	if limitdays == 0 {
		return us.bprr.GetRecordById(ctx, uid)
	}
	return us.bprr.GetRecordByIdLimitDays(ctx, uid, limitdays)
}

// 添加一条血压记录
func (us *patientBpRecordService) AddById(ctx context.Context, uid uint, low, high, heartRate int) error {
	// log.Debug("添加血压记录信息", log.WithPair("uid", uid), log.WithPair("low", low), log.WithPair("high", high))
	return us.bprr.AddRecord(ctx, uid, low, high, heartRate)
}
