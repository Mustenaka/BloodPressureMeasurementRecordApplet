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

var _ repo.AdminUserRepo = (*adminUserRepo)(nil)

type adminUserRepo struct {
	ds db.IDataSource
}

// 创建一个新的AdminUserRepo
func NewAdminUserRepo(_ds db.IDataSource) *adminUserRepo {
	return &adminUserRepo{
		ds: _ds,
	}
}

// 通过用户名称获取管理员
func (ur *adminUserRepo) GetAdminUserByName(ctx context.Context, name string) (*model.AdminUser, error) {
	user := &model.AdminUser{}
	var count int64
	err := ur.ds.Master().Where("user_name = ?", name).Find(user).Count(&count).Error
	if count == 0 {
		err = errors.New("record not found")
	}
	return user, err
}

// 通过userid获取用户
func (ur *adminUserRepo) GetAdminUserById(ctx context.Context, id uint) (*model.AdminUser, error) {
	user := &model.AdminUser{}
	var count int64
	err := ur.ds.Master().Where("user_id = ?", id).Find(user).Count(&count).Error
	if count == 0 {
		err = errors.New("record not found")
	}
	return user, err
}

// 通过用户米、密码创建新用户（管理端用户使用）
func (ur *adminUserRepo) AddAdminUserByNamePassword(ctx context.Context, name, password string) error {
	user := &model.AdminUser{
		AdminName: name,
		Password:  password,
	}
	err := ur.ds.Master().Create(user).Error
	return err
}

// 更新用户基本信息
func (ur *adminUserRepo) UpdateAdminUserDetail(ctx context.Context, srcUser *model.AdminUser, realname, telephone, email, brithday, sex string) error {
	nowTime := timeconvert.NowDateTimeString()
	result := ur.ds.Master().Where(&model.AdminUser{UserId: srcUser.UserId}).Model(&srcUser).Updates(&model.AdminUser{
		RealName: realname,
		Tel:      strtools.UpdateNotNullStirng(telephone, srcUser.Tel),
		Email:    strtools.UpdateNotNullStirng(email, srcUser.Email),
		Birthday: strtools.UpdateNotNullStirng(brithday, srcUser.Birthday),
		Sex:      sex,
		LastTime: nowTime,
	})
	return result.Error
}

// 更新用户密码
func (ur *adminUserRepo) UpdateAdminUserPassword(ctx context.Context, srcUser *model.AdminUser, password string) error {
	nowTime := timeconvert.NowDateTimeString()
	result := ur.ds.Master().Where(&model.AdminUser{UserId: srcUser.UserId}).Model(&srcUser).Updates(&model.AdminUser{
		Password: password,
		LastTime: nowTime,
	})
	return result.Error
}

// 解冻、冻结用户密码(更新用户状态)
func (ur *adminUserRepo) UpdateAdminUserStatus(ctx context.Context, id uint, status string) error {
	// 状态检查、日志记录、错误提前返回
	switch status {
	case "开启":
		log.Info("开启用户", log.WithPair("id", id))
	case "关闭":
		log.Info("关闭用户", log.WithPair("id", id))
	default:
		return errors.New("status string error")
	}

	// 更新状态
	result := ur.ds.Master().Where(&model.AdminUser{UserId: id}).Model(&model.AdminUser{}).Updates(&model.AdminUser{
		Status: status,
	})
	return result.Error
}
