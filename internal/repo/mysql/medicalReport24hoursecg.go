package mysql

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"BloodPressure/pkg/db"
	"BloodPressure/pkg/log"
	timeconvert "BloodPressure/tools/timeConvert"
	"context"
)

var _ repo.MedicalReport24hoursecgRepo = (*medicalReport24hoursecgRepo)(nil)

type medicalReport24hoursecgRepo struct {
	ds db.IDataSource
}

// 创建一个新的UserRepo
func NewMedicalReport24hoursecgRepo(_ds db.IDataSource) *medicalReport24hoursecgRepo {
	return &medicalReport24hoursecgRepo{
		ds: _ds,
	}
}

// 添加24小时动态ECG
func (ur *medicalReport24hoursecgRepo) AddEcg(ctx context.Context, id uint, data string, average int) error {
	createDatetime := timeconvert.NowDateTimeString()
	ecgDetails := &model.MedicalReport24hoursecg{
		UserId:   id,
		Data:     data,
		Average:  average,
		CreateAt: createDatetime,
	}
	log.Debug("添加24小时动态ECG", log.WithPair("ECG", ecgDetails.Data))
	err := ur.ds.Master().Create(ecgDetails).Error
	return err
}

// 自定义时间添加动态ECG
func (ur *medicalReport24hoursecgRepo) AddEcgWithTime(ctx context.Context, id uint, data string, average int, createAt string) error {
	ecgDetails := &model.MedicalReport24hoursecg{
		UserId:   id,
		Data:     data,
		Average:  average,
		CreateAt: createAt,
	}
	log.Debug("添加24小时动态ECG", log.WithPair("ECG", ecgDetails.Data))
	err := ur.ds.Master().Create(ecgDetails).Error
	return err
}

// 查询24小时动态ECG
func (ur *medicalReport24hoursecgRepo) GetEcgById(ctx context.Context, id uint) ([]model.MedicalReport24hoursecg, error) {
	ecgDetails := []model.MedicalReport24hoursecg{}
	err := ur.ds.Master().Where("user_id = ?", id).Find(&ecgDetails).Error
	return ecgDetails, err
}

// 通过limit查询限制的24小时动态ECG
func (ur *medicalReport24hoursecgRepo) GetEcgByIdLimit(ctx context.Context, id uint, limit int) ([]model.MedicalReport24hoursecg, error) {
	ecgDetails := []model.MedicalReport24hoursecg{}
	err := ur.ds.Master().Where("user_id = ?", id).Limit(limit).Find(&ecgDetails).Error
	return ecgDetails, err
}

// 删除该用户全部的24小时动态ECG
func (ur *medicalReport24hoursecgRepo) DeleteEcgByUserId(ctx context.Context, userId uint) error {
	err := ur.ds.Master().Where("user_id = ?", userId).Delete(&model.MedicalReport24hoursecg{}).Error
	return err
}

// 通过id删除详细的24小时动态ECG
func (ur *medicalReport24hoursecgRepo) DeleteEcgById(ctx context.Context, id uint) error {
	err := ur.ds.Master().Where("id = ?", id).Delete(&model.MedicalReport24hoursecg{}).Error
	return err
}
