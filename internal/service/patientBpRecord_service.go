package service

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"BloodPressure/pkg/log"
	"context"
)

var _ PatientBpRecordService = (*patientBpRecordService)(nil)

// PatientBpRecordService 定义用户操作服务接口
type PatientBpRecordService interface {
	// 查询
	GetById(ctx context.Context, uid uint) ([]model.PatientBpRecord, error)
	// 添加一条记录
	AddById(ctx context.Context, uid uint, low, high int) error
}

// patientBpRecordService 实现UserService接口
type patientBpRecordService struct {
	ur repo.PatientBpRecordRepo
}

// 新血压记录服务
func NewPBPRecordService(_ur repo.PatientBpRecordRepo) *patientBpRecordService {
	return &patientBpRecordService{
		ur: _ur,
	}
}

// 获取全部的测量数据
func (us *patientBpRecordService) GetById(ctx context.Context, uid uint) ([]model.PatientBpRecord, error) {
	return us.ur.GetRecordById(ctx, uid)
}

// 添加一条血压记录
func (us *patientBpRecordService) AddById(ctx context.Context, uid uint, low, high int) error {
	log.Debug("添加血压记录信息", log.WithPair("uid", uid), log.WithPair("low", low), log.WithPair("high", high))
	return us.ur.AddRecord(ctx, uid, low, high)
}
