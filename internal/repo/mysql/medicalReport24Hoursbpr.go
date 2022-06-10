package mysql

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"BloodPressure/pkg/db"
	"context"
)

var _ repo.MedicalReport24HoursbprRepo = (*medicalReport24HoursbprRepo)(nil)

type medicalReport24HoursbprRepo struct {
	ds db.IDataSource
}

// 创建一个新的UserRepo
func NewMedicalReport24HoursbprRepo(_ds db.IDataSource) *medicalReport24HoursbprRepo {
	return &medicalReport24HoursbprRepo{
		ds: _ds,
	}
}

// 添加24小时动态血压报告
func (ur *medicalReport24HoursbprRepo) Addbpr(ctx context.Context, id uint, report *model.MedicalReport24hoursbpr) error {
	// 添加 MedicalReport24hoursbpr信息
	err := ur.ds.Master().Create(report).Error
	return err
}

// 获取动态血压报告
func (ur *medicalReport24HoursbprRepo) GetbprById(ctx context.Context, id uint) ([]model.MedicalReport24hoursbpr, error) {
	var report []model.MedicalReport24hoursbpr
	err := ur.ds.Master().Where("user_id = ?", id).Find(&report).Error
	return report, err
}

// 通过limit查询限制的动态血压报告
func (ur *medicalReport24HoursbprRepo) GetbprByIdLimit(ctx context.Context, id uint, limit int) ([]model.MedicalReport24hoursbpr, error) {
	var report []model.MedicalReport24hoursbpr
	err := ur.ds.Master().Where("user_id = ?", id).Limit(limit).Find(&report).Error
	return report, err
}

// 删除该用户的全部动态血压报告
func (ur *medicalReport24HoursbprRepo) DeletebprByUserId(ctx context.Context, userId uint) error {
	err := ur.ds.Master().Where("user_id = ?", userId).Delete(&model.MedicalReport24hoursbpr{}).Error
	return err
}

// 通过id删除动态血压报告
func (ur *medicalReport24HoursbprRepo) DeletebprById(ctx context.Context, id uint) error {
	err := ur.ds.Master().Where("id = ?", id).Delete(&model.MedicalReport24hoursbpr{}).Error
	return err
}
