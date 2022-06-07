package baseuser

import (
	"BloodPressure/pkg/constant"
	"BloodPressure/pkg/errors"
	"BloodPressure/pkg/errors/code"
	"BloodPressure/pkg/response"
	"BloodPressure/tools/security"
	"context"

	"github.com/gin-gonic/gin"
)

// 修改密码
func (uh *BaseUserHandler) UpdateUserPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义基本结构
		type RegisterParam struct {
			Password string `json:"password" binding:"required"`
		}

		// 检验基本结构
		var param RegisterParam
		if err := c.ShouldBind(&param); err != nil {
			response.JSON(c, errors.Wrap(err, code.ValidateErr, "存在必要信息未填写"), nil)
			return
		}

		// 通过id找到基本用户
		uid := c.GetUint(constant.UserID)
		baseUser, err := uh.userSrv.GetById(context.TODO(), uid)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.NotFoundErr, "用户信息为空"), nil)
		}

		// 对密码进行MD5加密
		securityPassword := security.Md5(param.Password)

		// 修改用户信息
		err = uh.userSrv.UpdatePassword(context.TODO(), baseUser, securityPassword)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.UserUpdateErr, "用户信息修改失败"), nil)
			return
		}

		// 返回这个结果
		response.JSON(c, nil, struct {
			Result string `json:"result"`
		}{
			Result: "update successful",
		})
	}
}

// 修改用户信息
func (uh *BaseUserHandler) UpdateUserDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义基本结构
		type RegisterParam struct {
			RealName  string `json:"realname"`
			UserName  string `json:"username"`
			Telephone string `json:"telephone"`
			Email     string `json:"email"`
			Brithday  string `json:"brithday"`
			Sex       string `json:"sex"`
			AvatarUrl string `json:"avatarUrl"`
		}

		// 检验基本结构
		var param RegisterParam
		if err := c.ShouldBind(&param); err != nil {
			response.JSON(c, errors.Wrap(err, code.ValidateErr, "存在必要信息未填写"), nil)
			return
		}

		// 通过id找到基本用户
		uid := c.GetUint(constant.UserID)
		baseUser, err := uh.userSrv.GetById(context.TODO(), uid)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.NotFoundErr, "用户信息为空"), nil)
			return
		}

		// 修改用户信息
		err = uh.userSrv.UpdateDetail(context.TODO(),
			baseUser,
			param.RealName,
			param.UserName,
			param.Telephone,
			param.Email,
			param.Brithday,
			param.Sex,
			param.AvatarUrl)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.UserUpdateErr, "用户信息修改失败"), nil)
			return
		}

		// 返回这个结果
		response.JSON(c, nil, struct {
			Result string `json:"result"`
		}{
			Result: "update successful",
		})
	}
}
