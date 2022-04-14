package mysql

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"BloodPressure/pkg/db"
	"context"
	"errors"
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
func (ur *patientInfoRepo) AddInfo(ctx context.Context, patientInfo *model.PatientInfo) error {
	// 添加 Patientinfo信息
	err := ur.ds.Master().Create(patientInfo).Error
	return err
}

// 获取记录
func (ur *patientInfoRepo) GetInfoById(ctx context.Context, id uint) (*model.PatientInfo, error) {
	patientinfo := &model.PatientInfo{}
	var count int64
	// err := ur.ds.Master().Where("user_id = ?", id).Find(patientinfo).Count(&count).Error
	err := ur.ds.Master().Where(&model.PatientInfo{UserId: id}).Find(patientinfo).Count(&count).Error
	if count == 0 {
		err = errors.New("record not found")
	}
	return patientinfo, err
}

// 更新记录
func (ur *patientInfoRepo) UpdateInfoById(ctx context.Context, id uint, patientInfo *model.PatientInfo) error {
	err := ur.ds.Master().Where(&model.PatientInfo{UserId: id}).Updates(patientInfo).Error
	return err
}

// 删除记录
func (ur *patientInfoRepo) DeleteInfoByID(ctx context.Context, id uint) error {
	err := ur.ds.Master().Where(&model.PatientInfo{UserId: id}).Delete(&model.PatientInfo{}).Error
	return err
}
