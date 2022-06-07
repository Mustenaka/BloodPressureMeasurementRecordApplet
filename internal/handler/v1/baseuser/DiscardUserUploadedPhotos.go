package baseuser

import (
	"BloodPressure/pkg/constant"
	"BloodPressure/pkg/errors"
	"BloodPressure/pkg/errors/code"
	"BloodPressure/pkg/log"
	"BloodPressure/pkg/response"
	"context"

	"github.com/gin-gonic/gin"
)

// 抛弃用户上传的照片
func (uh *BaseUserHandler) DiscardUserUploadedPhotos() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.GetUint(constant.UserID)
		// 通过id找到基本用户
		baseUser, err := uh.userSrv.GetById(context.TODO(), uid)

		if err != nil {
			response.JSON(c, errors.Wrap(err, code.NotFoundErr, "用户信息为空"), nil)
		} else {
			response.JSON(c, nil, baseUser)
		}

		// 获取用户上传照片，解析到表单
		form, err := c.MultipartForm()
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.BadRequestErr, "获取用户上传照片失败"), nil)
		}

		// 从表单中找到所有的文件
		files := form.File["pictures"]
		for _, file := range files {
			// 获取文件名
			filename := file.Filename
			// 获取文件内容
			log.Info("文件获取", log.WithPair("filename", filename))
		}

		response.JSON(c, nil, struct {
			Message string `json:"message"`
		}{
			Message: "用户上传的照片成功",
		})
	}
}
