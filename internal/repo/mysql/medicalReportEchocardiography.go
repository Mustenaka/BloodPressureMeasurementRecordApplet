package mysql

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"BloodPressure/pkg/db"
	"BloodPressure/pkg/log"
	timeconvert "BloodPressure/tools/timeConvert"
	"context"
)

var _ repo.MedicalReportEchocardiographyRepo = (*medicalReportEchocardiographyRepo)(nil)

type medicalReportEchocardiographyRepo struct {
	ds db.IDataSource
}

// 创建一个新的UserRepo
func NewMedicalReportEchocardiographyRepo(_ds db.IDataSource) *medicalReportEchocardiographyRepo {
	return &medicalReportEchocardiographyRepo{
		ds: _ds,
	}
}

//  添加心超
func (ur *medicalReportEchocardiographyRepo) AddEchocardiography(ctx context.Context, id uint, ef, lvidd, lvids int) error {
	createDatetime := timeconvert.NowDateTimeString()
	tongueDetails := &model.MedicalReportEchocardiography{
		UserId:   id,
		Ef:       ef,
		Lvidd:    lvidd,
		Lvids:    lvids,
		CreateAt: createDatetime,
	}
	log.Debug("添加心超", log.WithPair("心超", tongueDetails.Ef))
	err := ur.ds.Master().Create(tongueDetails).Error
	return err
}

// 自定义时间添加心超
func (ur *medicalReportEchocardiographyRepo) AddEchocardiographyWithTime(ctx context.Context, id uint, ef, lvidd, lvids int, createAt string) error {
	tongueDetails := &model.MedicalReportEchocardiography{
		UserId:   id,
		Ef:       ef,
		Lvidd:    lvidd,
		Lvids:    lvids,
		CreateAt: createAt,
	}
	log.Debug("添加心超", log.WithPair("心超", tongueDetails.Ef))
	err := ur.ds.Master().Create(tongueDetails).Error
	return err
}

// 获取心超
func (ur *medicalReportEchocardiographyRepo) GetEchocardiographyById(ctx context.Context, id uint) ([]model.MedicalReportEchocardiography, error) {
	tongueDetails := []model.MedicalReportEchocardiography{}
	err := ur.ds.Master().Where("user_id = ?", id).Find(&tongueDetails).Error
	return tongueDetails, err
}

// 通过limit查询限制的心超
func (ur *medicalReportEchocardiographyRepo) GetEchocardiographyByIdLimit(ctx context.Context, id uint, limit int) ([]model.MedicalReportEchocardiography, error) {
	tongueDetails := []model.MedicalReportEchocardiography{}
	err := ur.ds.Master().Where("user_id = ?", id).Limit(limit).Find(&tongueDetails).Error
	return tongueDetails, err
}

// 删除该用户全部心超信息
func (ur *medicalReportEchocardiographyRepo) DeleteEchocardiographyByUserId(ctx context.Context, userId uint) error {
	err := ur.ds.Master().Where("user_id = ?", userId).Delete(&model.MedicalReportEchocardiography{}).Error
	return err
}

// 通过id删除详细信息
func (ur *medicalReportEchocardiographyRepo) DeleteEchocardiographyById(ctx context.Context, id uint) error {
	err := ur.ds.Master().Where("id = ?", id).Delete(&model.MedicalReportEchocardiography{}).Error
	return err
}
