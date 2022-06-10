package mysql

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"BloodPressure/pkg/db"
	"BloodPressure/pkg/log"
	timeconvert "BloodPressure/tools/timeConvert"
	"context"
)

var _ repo.TongueDetailRepo = (*tongueDetailRepo)(nil)

type tongueDetailRepo struct {
	ds db.IDataSource
}

// 创建一个新的UserRepo
func NewTongueDetailRepo(_ds db.IDataSource) *tongueDetailRepo {
	return &tongueDetailRepo{
		ds: _ds,
	}
}

// 添加详细信息
func (ur *tongueDetailRepo) AddTongue(ctx context.Context, id uint, tongue, tongueCoating, pulse string) error {
	createDatetime := timeconvert.NowDateTimeString()
	tongueDetails := &model.TongueDetail{
		UserId:        id,
		Tongue:        tongue,
		TongueCoating: tongueCoating,
		Pulse:         pulse,
		CreateAt:      createDatetime,
	}
	log.Debug("添加详细信息", log.WithPair("详细信息", tongueDetails.Tongue))
	err := ur.ds.Master().Create(tongueDetails).Error
	return err
}

// 自定义时间添加详细信息
func (ur *tongueDetailRepo) AddTongueWithTime(ctx context.Context, id uint, tongue, tongueCoating, pulse, createAt string) error {
	tongueDetails := &model.TongueDetail{
		UserId:        id,
		Tongue:        tongue,
		TongueCoating: tongueCoating,
		Pulse:         pulse,
		CreateAt:      createAt,
	}
	log.Debug("添加详细信息", log.WithPair("详细信息", tongueDetails.Tongue))
	err := ur.ds.Master().Create(tongueDetails).Error
	return err
}

// 查询详细信息
func (ur *tongueDetailRepo) GetTongueById(ctx context.Context, id uint) ([]model.TongueDetail, error) {
	tongueDetails := []model.TongueDetail{}
	err := ur.ds.Master().Where("user_id = ?", id).Find(&tongueDetails).Error
	return tongueDetails, err
}

// 通过limit查询限制的详细信息
func (ur *tongueDetailRepo) GetTongueByIdLimit(ctx context.Context, id uint, limit int) ([]model.TongueDetail, error) {
	tongueDetails := []model.TongueDetail{}
	err := ur.ds.Master().Where("user_id = ?", id).Limit(limit).Find(&tongueDetails).Error
	return tongueDetails, err
}

// 删除该用户全部详细信息
func (ur *tongueDetailRepo) DeleteTongueByUserId(ctx context.Context, userId uint) error {
	err := ur.ds.Master().Where("user_id = ?", userId).Delete(&model.TongueDetail{}).Error
	return err
}

// 通过id删除详细信息
func (ur *tongueDetailRepo) DeleteTongueById(ctx context.Context, id uint) error {
	err := ur.ds.Master().Where("id = ?", id).Delete(&model.TongueDetail{}).Error
	return err
}
