package mysql

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"BloodPressure/pkg/db"
	"BloodPressure/pkg/log"
	timeconvert "BloodPressure/tools/timeConvert"
	"context"
)

var _ repo.TestIndicatorBnpRepo = (*testIndicatorBnpRepo)(nil)

type testIndicatorBnpRepo struct {
	ds db.IDataSource
}

// 创建一个新的UserRepo
func NewTestIndicatorBnpRepo(_ds db.IDataSource) *testIndicatorBnpRepo {
	return &testIndicatorBnpRepo{
		ds: _ds,
	}
}

// 添加Bnp
func (ur *testIndicatorBnpRepo) AddBnp(ctx context.Context, id uint, data int) error {
	createDatetime := timeconvert.NowDateTimeString()
	tongueDetails := &model.TestIndicatorBnp{
		UserId:   id,
		Data:     data,
		CreateAt: createDatetime,
	}
	log.Debug("添加Bnp", log.WithPair("Bnp", tongueDetails.Data))
	err := ur.ds.Master().Create(tongueDetails).Error
	return err
}

// 自定义时间添加Bnp
func (ur *testIndicatorBnpRepo) AddBnpWithTime(ctx context.Context, id uint, data int, createAt string) error {
	tongueDetails := &model.TestIndicatorBnp{
		UserId:   id,
		Data:     data,
		CreateAt: createAt,
	}
	log.Debug("添加Bnp", log.WithPair("Bnp", tongueDetails.Data))
	err := ur.ds.Master().Create(tongueDetails).Error
	return err
}

// 获取Bnp
func (ur *testIndicatorBnpRepo) GetBnpById(ctx context.Context, id uint) ([]model.TestIndicatorBnp, error) {
	tongueDetails := []model.TestIndicatorBnp{}
	err := ur.ds.Master().Where("user_id = ?", id).Find(&tongueDetails).Error
	return tongueDetails, err
}

// 通过limit查询限制的Bnp
func (ur *testIndicatorBnpRepo) GetBnpByIdLimit(ctx context.Context, id uint, limit int) ([]model.TestIndicatorBnp, error) {
	tongueDetails := []model.TestIndicatorBnp{}
	err := ur.ds.Master().Where("user_id = ?", id).Limit(limit).Find(&tongueDetails).Error
	return tongueDetails, err
}

// 删除该用户全部Bnp信息
func (ur *testIndicatorBnpRepo) DeleteBnpByUserId(ctx context.Context, userId uint) error {
	err := ur.ds.Master().Where("user_id = ?", userId).Delete(&model.TestIndicatorBnp{}).Error
	return err
}

// 通过id删除详细信息
func (ur *testIndicatorBnpRepo) DeleteBnpById(ctx context.Context, id uint) error {
	err := ur.ds.Master().Where("id = ?", id).Delete(&model.TestIndicatorBnp{}).Error
	return err
}
