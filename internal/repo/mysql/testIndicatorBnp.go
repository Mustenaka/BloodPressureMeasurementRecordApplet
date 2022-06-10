package mysql

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"BloodPressure/pkg/db"
	"BloodPressure/pkg/log"
	timeconvert "BloodPressure/tools/timeConvert"
	"context"
)

var _ repo.TestIndicatorBnp = (*testIndicatorBnp)(nil)

type testIndicatorBnp struct {
	ds db.IDataSource
}

// 创建一个新的UserRepo
func NewTestIndicatorBnp(_ds db.IDataSource) *testIndicatorBnp {
	return &testIndicatorBnp{
		ds: _ds,
	}
}

// 添加Bnp
func (ur *testIndicatorBnp) AddBnp(ctx context.Context, id uint, data int) error {
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

// 自定义事件添加Bnp
func (ur *testIndicatorBnp) AddBnpWithTime(ctx context.Context, id uint, data int, createAt string) error {
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
func (ur *testIndicatorBnp) GetBnpById(ctx context.Context, id uint) ([]model.TestIndicatorBnp, error) {
	tongueDetails := []model.TestIndicatorBnp{}
	err := ur.ds.Master().Where("user_id = ?", id).Find(&tongueDetails).Error
	return tongueDetails, err
}

// 通过limit查询限制的Bnp
func (ur *testIndicatorBnp) GetBnpByIdLimit(ctx context.Context, id uint, limit int) ([]model.TestIndicatorBnp, error) {
	tongueDetails := []model.TestIndicatorBnp{}
	err := ur.ds.Master().Where("user_id = ?", id).Limit(limit).Find(&tongueDetails).Error
	return tongueDetails, err
}

// 删除该用户全部Bnp信息
func (ur *testIndicatorBnp) DeleteBnpByUserId(ctx context.Context, userId uint) error {
	err := ur.ds.Master().Where("user_id = ?", userId).Delete(&model.TestIndicatorBnp{}).Error
	return err
}

// 通过id删除详细信息
func (ur *testIndicatorBnp) DeleteBnpById(ctx context.Context, id uint) error {
	err := ur.ds.Master().Where("id = ?", id).Delete(&model.TestIndicatorBnp{}).Error
	return err
}
