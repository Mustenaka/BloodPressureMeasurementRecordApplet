package mysql

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"BloodPressure/pkg/db"
	"context"
)

var _ repo.PatientInfoRepo = (*patientInfoRepo)(nil)

type patientInfoRepo struct {
	ds db.IDataSource
}

// 创建一个新的UserRepo
func NewPatientInfoRepo(_ds db.IDataSource) *patientInfoRepo {
	return &patientInfoRepo{
		ds: _ds,
	}
}

// 添加病历信息记录
func (ur *patientInfoRepo) AddInfo(ctx context.Context, id uint) error {
	return nil
}

// 获取记录组
func (ur *patientInfoRepo) GetInfoById(ctx context.Context, id uint) (*model.PatientInfo, error) {
	return nil, nil
}

// 更新记录
func (ur *patientInfoRepo) UpdateInfoById(ctx context.Context, id uint) (*model.PatientInfo, error) {
	return nil, nil
}
