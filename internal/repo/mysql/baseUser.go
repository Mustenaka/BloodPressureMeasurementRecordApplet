package mysql

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"BloodPressure/pkg/db"
	"BloodPressure/pkg/log"
	strtools "BloodPressure/tools/strTools"
	timeconvert "BloodPressure/tools/timeConvert"
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
	err := ur.ds.Master().Where("user_name = ?", name).Find(user).Count(&count).Error
	if count == 0 {
		err = errors.New("record not found")
	}
	return user, err
}

// 通过userid获取用户
func (ur *baseUserRepo) GetBaseUserById(ctx context.Context, uid uint) (*model.BaseUser, error) {
	user := &model.BaseUser{}
	var count int64
	err := ur.ds.Master().Where("user_id = ?", uid).Find(user).Count(&count).Error
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

// 通过用户名、密码创建新用户（内部接口、目前不需要不对外使用）
func (ur *baseUserRepo) AddBaseUserByNamePassword(ctx context.Context, name, password string) error {
	nowTime := timeconvert.NowDateTimeString()
	user := &model.BaseUser{
		UserName:   name,
		Password:   password,
		LastTime:   nowTime,
		CreateTime: nowTime,
		Status:     "开启",
	}
	err := ur.ds.Master().Create(user).Error
	return err
}

// 通过openid创建新用户（微信用户使用）
func (ur *baseUserRepo) AddBaseUserByDetail(ctx context.Context, name, openid, sex, avatarUrl string) error {
	nowTime := timeconvert.NowDateTimeString()
	user := &model.BaseUser{
		UserName:   name,
		OpenId:     openid,
		Sex:        sex,
		AvatarUrl:  avatarUrl, // 头像url
		LastTime:   nowTime,
		CreateTime: nowTime,
		Status:     "开启",
	}
	err := ur.ds.Master().Create(user).Error
	return err
}

// 更新用户基本信息 - 修改建议，srcUser改用传入id的方式, 然后tel,birthday,sex 采用可变参数
func (ur *baseUserRepo) UpdateBaseUserDetail(ctx context.Context, srcUser *model.BaseUser, realname, telephone, email, brithday, sex, avatarUrl string) error {
	nowTime := timeconvert.NowDateTimeString()
	result := ur.ds.Master().Where(&model.BaseUser{UserId: srcUser.UserId}).Model(&srcUser).Updates(&model.BaseUser{
		RealName:  realname,
		Tel:       strtools.UpdateNotNullStirng(telephone, srcUser.Tel),
		Email:     strtools.UpdateNotNullStirng(email, srcUser.Email),
		Birthday:  strtools.UpdateNotNullStirng(brithday, srcUser.Birthday),
		Sex:       sex,
		AvatarUrl: avatarUrl,
		LastTime:  nowTime,
	})
	log.Debug("Update db", log.WithPair("affect count", result.RowsAffected))
	return result.Error
}

// 更新用户密码
func (ur *baseUserRepo) UpdateBaseUserPassword(ctx context.Context, srcUser *model.BaseUser, password string) error {
	nowTime := timeconvert.NowDateTimeString()
	result := ur.ds.Master().Where(&model.BaseUser{UserId: srcUser.UserId}).Model(&srcUser).Updates(&model.BaseUser{
		Password: password,
		LastTime: nowTime,
	})
	log.Debug("Update db", log.WithPair("affect count", result.RowsAffected))
	return result.Error
}
