package baseuser

import (
	"BloodPressure/pkg/errors"
	"BloodPressure/pkg/errors/code"
	"BloodPressure/pkg/response"
	"BloodPressure/tools/security"
	"context"
	e "errors"

	"github.com/gin-gonic/gin"
)

// 管理员注册
func (uh *BaseUserHandler) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义基本结构
		type RegisterParam struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		// 检验基本结构
		var param RegisterParam
		if err := c.ShouldBind(&param); err != nil {
			response.JSON(c, errors.Wrap(err, code.ValidateErr, "存在必要信息未填写"), nil)
			return
		}

		// 查询用户是否存在
		_, err := uh.userSrv.GetByName(context.TODO(), param.Username)
		if err == nil {
			response.JSON(c, errors.Wrap(e.New("account repeated existence"), code.UserRegisterErr, "注册失败，用户已存在"), nil)
			return
		}

		// 对密码进行MD5加密
		securityPassword := security.Md5(param.Password)
		err = uh.userSrv.AddByNameAndPassword(context.TODO(), param.Username, securityPassword)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.UserRegisterErr, "注册失败，无法注册"), nil)
		}

		// 返回这个结果
		response.JSON(c, nil, struct {
			Result string `json:"result"`
		}{
			Result: "register successful",
		})
	}
}

// 用户注册
func (uh *BaseUserHandler) WeRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义基本结构
		type RegisterParam struct {
			Username  string `json:"username" binding:"required"`
			OpenId    string `json:"openid"   binding:"required"`
			RealName  string `json:"realname"   binding:"required"`
			Telephone string `json:"telephone"   binding:"required"`
			Email     string `json:"email"`
			Brithday  string `json:"brithday"`
			Sex       string `json:"sex"`
		}

		// 检验基本结构
		var param RegisterParam
		if err := c.ShouldBind(&param); err != nil {
			response.JSON(c, errors.Wrap(err, code.ValidateErr, "存在必要信息未填写"), nil)
			return
		}

		// 查询用户是否存在
		_, err := uh.userSrv.GetByOpenid(context.TODO(), param.OpenId)
		if err == nil {
			response.JSON(c, errors.Wrap(e.New("account repeated existence"), code.UserRegisterErr, "注册失败，用户已存在"), nil)
			return
		}

		// 注册信息
		err = uh.userSrv.AddByDetail(context.TODO(),
			param.Username,
			param.OpenId,
			param.RealName,
			param.Telephone,
			param.Email,
			param.Brithday,
			param.Sex)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.UserRegisterErr, "注册失败，无法注册"), nil)
		}

		// 返回这个结果
		response.JSON(c, nil, struct {
			Result string `json:"result"`
		}{
			Result: "wechat user register successful",
		})
	}
}
