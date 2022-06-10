package mysql

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"BloodPressure/pkg/db"
	"BloodPressure/pkg/log"
	timeconvert "BloodPressure/tools/timeConvert"
	"context"
)

var _ repo.MedicalReportEcg = (*medicalReportEcg)(nil)

type medicalReportEcg struct {
	ds db.IDataSource
}

// 创建一个新的UserRepo
func NewMedicalReportEcg(_ds db.IDataSource) *medicalReportEcg {
	return &medicalReportEcg{
		ds: _ds,
	}
}

// 添加ecg
func (ur *medicalReportEcg) AddEcg(ctx context.Context, id uint, data string) error {
	createDatetime := timeconvert.NowDateTimeString()
	ecgDetails := &model.MedicalReportEcg{
		UserId:   id,
		Data:     data,
		CreateAt: createDatetime,
	}
	log.Debug("添加ecg", log.WithPair("ecg", ecgDetails.Data))
	err := ur.ds.Master().Create(ecgDetails).Error
	return err
}

// 自定义时间添加ecg
func (ur *medicalReportEcg) AddEcgWithTime(ctx context.Context, id uint, data string, createAt string) error {
	ecgDetails := &model.MedicalReportEcg{
		UserId:   id,
		Data:     data,
		CreateAt: createAt,
	}
	log.Debug("添加ecg", log.WithPair("ecg", ecgDetails.Data))
	err := ur.ds.Master().Create(ecgDetails).Error
	return err
}

// 获取ecg
func (ur *medicalReportEcg) GetEcgById(ctx context.Context, id uint) ([]model.MedicalReportEcg, error) {
	ecgDetails := []model.MedicalReportEcg{}
	err := ur.ds.Master().Where("user_id = ?", id).Find(&ecgDetails).Error
	return ecgDetails, err
}

// 通过limit查询限制的ecg
func (ur *medicalReportEcg) GetEcgByIdLimit(ctx context.Context, id uint, limit int) ([]model.MedicalReportEcg, error) {
	ecgDetails := []model.MedicalReportEcg{}
	err := ur.ds.Master().Where("user_id = ?", id).Limit(limit).Find(&ecgDetails).Error
	return ecgDetails, err
}

// 删除该用户全部ecg信息
func (ur *medicalReportEcg) DeleteEcgById(ctx context.Context, id uint) error {
	err := ur.ds.Master().Where("user_id = ?", id).Delete(&model.MedicalReportEcg{}).Error
	return err
}

// 通过id删除详细信息
func (ur *medicalReportEcg) DeleteEcgByUserId(ctx context.Context, userId uint) error {
	err := ur.ds.Master().Where("user_id = ?", userId).Delete(&model.MedicalReportEcg{}).Error
	return err
}
