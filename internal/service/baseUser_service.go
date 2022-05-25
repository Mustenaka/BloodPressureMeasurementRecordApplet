package service

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"BloodPressure/pkg/errors"
	"BloodPressure/pkg/errors/code"
	"BloodPressure/tools/reg"
	"context"
)

var _ BaseUserService = (*baseUserService)(nil)

// BaseUserService 定义用户操作服务接口
type BaseUserService interface {
	// 查询
	GetByName(ctx context.Context, name string) (*model.BaseUser, error)
	GetById(ctx context.Context, uid uint) (*model.BaseUser, error)
	GetByOpenid(ctx context.Context, openid string) (*model.BaseUser, error)
	// 增加
	AddByNameAndPassword(ctx context.Context, name, password string) error
	AddByDetail(ctx context.Context, name, openid, sex, avatarUrl string) error
	// 修改
	UpdateDetail(ctx context.Context, src *model.BaseUser, realname, telephone, email, brithday, sex, avatarUrl string) error
	UpdatePassword(ctx context.Context, src *model.BaseUser, password string) error
	// 删除
}

// baseUserService 实现UserService接口
type baseUserService struct {
	ur repo.BaseUserRepo
}

// 新用户服务
func NewBaseUserService(_ur repo.BaseUserRepo) *baseUserService {
	return &baseUserService{
		ur: _ur,
	}
}

// GetByName 通过用户名 查找用户
func (us *baseUserService) GetByName(ctx context.Context, name string) (*model.BaseUser, error) {
	if len(name) == 0 {
		return nil, errors.WithCode(code.ValidateErr, "用户名称不能为空")
	}
	return us.ur.GetBaseUserByName(ctx, name)
}

// GetById 根据用户ID查找用户
func (us *baseUserService) GetById(ctx context.Context, uid uint) (*model.BaseUser, error) {
	return us.ur.GetBaseUserById(ctx, uid)
}

// GetByOpenid 通过openid找到目标用户
func (us *baseUserService) GetByOpenid(ctx context.Context, openid string) (*model.BaseUser, error) {
	if len(openid) == 0 {
		return nil, errors.WithCode(code.ValidateErr, "openid不能为空")
	}
	return us.ur.GetBaseUserByOpenId(ctx, openid)
}

// 通过用户名密码添加新用户（后台管理端使用）
func (us *baseUserService) AddByNameAndPassword(ctx context.Context, name, password string) error {
	return us.ur.AddBaseUserByNamePassword(ctx, name, password)
}

// 通过详细信息创建用户（微信小程序端口使用）
func (us *baseUserService) AddByDetail(ctx context.Context, name, openid, sex, avatarUrl string) error {
	return us.ur.AddBaseUserByDetail(ctx, name, openid, sex, avatarUrl)
}

// 更新用户详细信息
func (us *baseUserService) UpdateDetail(ctx context.Context, src *model.BaseUser, realname, telephone, email, brithday, sex, avatarUrl string) error {
	// 验证生日字段
	if brithday != "" && !reg.VerifyDateFormat(brithday) {
		return errors.WithCode(code.ValidateErr, "生日日期格式不对，正确 YYYY-MM-DD")
	}
	// 验证电话号码
	if telephone != "" && !reg.VerifyMobileFormat(telephone) {
		return errors.WithCode(code.ValidateErr, "电话号码格式错误")
	}
	// 验证邮箱
	if telephone != "" && !reg.VerifyEmailFormat(email) {
		return errors.WithCode(code.ValidateErr, "邮箱格式错误")
	}

	return us.ur.UpdateBaseUserDetail(ctx, src, realname, telephone, email, brithday, sex, avatarUrl)
}

// 更新用户密码
func (us *baseUserService) UpdatePassword(ctx context.Context, src *model.BaseUser, password string) error {
	return us.ur.UpdateBaseUserPassword(ctx, src, password)
}
