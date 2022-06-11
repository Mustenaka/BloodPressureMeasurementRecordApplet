package mysql

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"BloodPressure/pkg/db"
	"BloodPressure/pkg/log"
	timeconvert "BloodPressure/tools/timeConvert"
	"context"
)

var _ repo.TestIndicatorCreatinineRepo = (*testIndicatorCreatinineRepo)(nil)

type testIndicatorCreatinineRepo struct {
	ds db.IDataSource
}

// 创建一个新的UserRepo
func NewTestIndicatorCreatinine(_ds db.IDataSource) *testIndicatorCreatinineRepo {
	return &testIndicatorCreatinineRepo{
		ds: _ds,
	}
}

// 添加肌酐
func (ur *testIndicatorCreatinineRepo) AddBnp(ctx context.Context, id uint, data int) error {
	createDatetime := timeconvert.NowDateTimeString()
	tongueDetails := &model.TestIndicatorCreatinine{
		UserId:   id,
		Data:     data,
		CreateAt: createDatetime,
	}
	log.Debug("添加肌酐", log.WithPair("肌酐", tongueDetails.Data))
	err := ur.ds.Master().Create(tongueDetails).Error
	return err
}

// 自定义事件添加肌酐
func (ur *testIndicatorCreatinineRepo) AddBnpWithTime(ctx context.Context, id uint, data int, createAt string) error {
	tongueDetails := &model.TestIndicatorCreatinine{
		UserId:   id,
		Data:     data,
		CreateAt: createAt,
	}
	log.Debug("添加肌酐", log.WithPair("肌酐", tongueDetails.Data))
	err := ur.ds.Master().Create(tongueDetails).Error
	return err
}

// 获取肌酐
func (ur *testIndicatorCreatinineRepo) GetBnpById(ctx context.Context, id uint) ([]model.TestIndicatorCreatinine, error) {
	tongueDetails := []model.TestIndicatorCreatinine{}
	err := ur.ds.Master().Where("user_id = ?", id).Find(&tongueDetails).Error
	return tongueDetails, err
}

// 通过limit查询限制的肌酐
func (ur *testIndicatorCreatinineRepo) GetBnpByIdLimit(ctx context.Context, id uint, limit int) ([]model.TestIndicatorCreatinine, error) {
	tongueDetails := []model.TestIndicatorCreatinine{}
	err := ur.ds.Master().Where("user_id = ?", id).Limit(limit).Find(&tongueDetails).Error
	return tongueDetails, err
}

// 删除该用户的全部肌酐信息
func (ur *testIndicatorCreatinineRepo) DeleteBnpByUserId(ctx context.Context, userId uint) error {
	err := ur.ds.Master().Where("user_id = ?", userId).Delete(&model.TestIndicatorCreatinine{}).Error
	return err
}

// 通过id删除肌酐信息
func (ur *testIndicatorCreatinineRepo) DeleteBnpById(ctx context.Context, id uint) error {
	err := ur.ds.Master().Where("id = ?", id).Delete(&model.TestIndicatorCreatinine{}).Error
	return err
}
