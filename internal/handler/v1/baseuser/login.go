package baseuser

import (
	"BloodPressure/pkg/config"
	"BloodPressure/pkg/errors"
	"BloodPressure/pkg/errors/code"
	"BloodPressure/pkg/jwt"
	"BloodPressure/pkg/response"
	jtime "BloodPressure/pkg/time"
	"BloodPressure/tools/security"
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

// 登录模块 - 通过用户名，密码登录
func (uh *BaseUserHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		type LoginParam struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
			// OpenId   string `json:"openid"   binding:"required"`
		}

		var param LoginParam
		if err := c.ShouldBind(&param); err != nil {
			response.JSON(c, errors.Wrap(err, code.ValidateErr, "用户名和密码不能为空"), nil)
			return
		}

		// 查询用户信息
		user, err := uh.userSrv.GetByName(context.TODO(), param.Username)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.UserLoginErr, "登录失败，用户不存在"), nil)
			return
		}

		// 检查用户名密码
		if !security.ValidatePassword(param.Password, user.Password) {
			response.JSON(c, errors.WithCode(code.UserLoginErr, "登录失败，用户名、密码不匹配"), nil)
			return
		}

		// 生成jwt token
		expireAt := time.Now().Add(24 * 7 * time.Hour)                                 // 当前时间+7天
		claims := jwt.BuildClaims(expireAt, user.UserId)                               // 与当前用户id生成一个声明
		token, err := jwt.GenToken(claims, config.GlobalConfig.ServerConfig.JwtSecret) // 再通过jwtsecret生成token

		// 生成token失败了
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.UserLoginErr, "生成用户授权token失败"), nil)
			return
		}

		// 返回这个token
		response.JSON(c, nil, struct {
			Token    string         `json:"token"`
			ExpireAt jtime.JsonTime `json:"expire_at"`
		}{
			Token:    token,
			ExpireAt: jtime.JsonTime(expireAt),
		})
	}
}

// 登录模块 - 直接通过openid做到一键登录
func (uh *BaseUserHandler) LoginWithOpenid() gin.HandlerFunc {
	return func(c *gin.Context) {
		type LoginParam struct {
			OpenId string `json:"openid"   binding:"required"`
		}

		var param LoginParam
		if err := c.ShouldBind(&param); err != nil {
			response.JSON(c, errors.Wrap(err, code.ValidateErr, "openid 不能为空"), nil)
			return
		}

		// 查询用户信息
		user, err := uh.userSrv.GetByOpenid(context.TODO(), param.OpenId)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.UserLoginErr, "登录失败，用户不存在"), nil)
			return
		}

		// 生成jwt token
		expireAt := time.Now().Add(24 * 7 * time.Hour)                                 // 当前时间+7天
		claims := jwt.BuildClaims(expireAt, user.UserId)                               // 与当前用户id生成一个声明
		token, err := jwt.GenToken(claims, config.GlobalConfig.ServerConfig.JwtSecret) // 再通过jwtsecret生成token

		// 生成token失败了
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.UserLoginErr, "生成用户授权token失败"), nil)
			return
		}

		// 返回这个token
		response.JSON(c, nil, struct {
			Token    string         `json:"token"`
			ExpireAt jtime.JsonTime `json:"expire_at"`
		}{
			Token:    token,
			ExpireAt: jtime.JsonTime(expireAt),
		})
	}
}
