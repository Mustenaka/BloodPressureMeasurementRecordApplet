package mysql

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"BloodPressure/pkg/db"
	"context"
	"errors"
)

var _ repo.BaseUserRepo = (*baseUserRepo)(nil)

type baseUserRepo struct {
	ds db.IDataSource
}

// 创建一个新的UserRepo
func NewBaseUserRepo(_ds db.IDataSource) *baseUserRepo {
	return &baseUserRepo{
		ds: _ds,
	}
}

// 通过用户名称获取用户
func (ur *baseUserRepo) GetBaseUserByName(ctx context.Context, name string) (*model.BaseUser, error) {
	user := &model.BaseUser{}
	var count int64
	err := ur.ds.Master().Where("user_name = ?", name).Find(user).Error
	if count == 0 {
		err = errors.New("record not found")
	}
	return user, err
}

// 通过userid获取用户
func (ur *baseUserRepo) GetBaseUserById(ctx context.Context, uid uint) (*model.BaseUser, error) {
	user := &model.BaseUser{}
	var count int64
	err := ur.ds.Master().Where("user_id = ?", uid).Find(user).Error
	if count == 0 {
		err = errors.New("record not found")
	}
	return user, err
}

// 通过openid获取用户信息
func (ur *baseUserRepo) GetBaseUserByOpenId(ctx context.Context, openid string) (*model.BaseUser, error) {
	user := &model.BaseUser{}
	var count int64
	err := ur.ds.Master().Where("open_id = ?", openid).Find(user).Count(&count).Error
	if count == 0 {
		err = errors.New("record not found")
	}
	return user, err
}

// 通过用户名、密码创建新用户
func (ur *baseUserRepo) AddBaseUserByAdmin(ctx context.Context, name, password string) (*model.BaseUser, error) {
	user := &model.BaseUser{
		UserName: name,
		Password: password,
	}
	err := ur.ds.Master().Create(user).Error
	return user, err
}

// 通过用户名、密码创建新用户
func (ur *baseUserRepo) AddBaseUserByDetail(ctx context.Context, name, openid, realname, telephone, email, brithday, sex string) (*model.BaseUser, error) {
	user := &model.BaseUser{
		UserName: name,
		OpenId:   openid,
		RealName: realname,
		Tel:      telephone,
		Email:    email,
		Birthday: brithday, // 时间之间转换成字符串给mysql接收，会根据字符串格式进行自动转换的
		Sex:      sex,
	}
	err := ur.ds.Master().Create(user).Error
	return user, err
}
