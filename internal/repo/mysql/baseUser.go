package mysql

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"BloodPressure/pkg/db"
	"context"
)

var _ repo.BaseUserRepo = (*baseUserRepo)(nil)

type baseUserRepo struct {
	ds db.IDataSource
}

func NewUserRepo(_ds db.IDataSource) *baseUserRepo {
	return &baseUserRepo{
		ds: _ds,
	}
}

func (ur *baseUserRepo) GetUserByName(ctx context.Context, name string) (*model.BaseUser, error) {
	user := &model.BaseUser{}
	err := ur.ds.Master().Where("user_name = ?", name).Find(user).Error
	return user, err
}

func (ur *baseUserRepo) GetUserById(ctx context.Context, uid uint) (*model.BaseUser, error) {
	user := &model.BaseUser{}
	err := ur.ds.Master().Where("id = ?", uid).Find(user).Error
	return user, err
}

func (ur *baseUserRepo) GetUserByOpenId(ctx context.Context, openid string) (*model.BaseUser, error) {
	user := &model.BaseUser{}
	err := ur.ds.Master().Where("open_id = ?", openid).Find(user).Error
	return user, err
}
