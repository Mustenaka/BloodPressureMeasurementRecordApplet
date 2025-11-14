package ping

import (
	eerrors "BloodPressure/pkg/errors"
	"BloodPressure/pkg/errors/code"
	"BloodPressure/pkg/response"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tiia "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tiia/v20190529"
)

// Ping ping服务器状态
func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "\r\nSUCCESS")
	}
}

// get image and send to tx sdk
func GetImageScore() gin.HandlerFunc {
	return func(c *gin.Context) {
		type ImageParam struct {
			Base64Image string `json:"image64"`
		}

		var param ImageParam
		if err := c.ShouldBind(&param); err != nil {
			response.JSON(c, eerrors.Wrap(err, code.ValidateErr, "error"), nil)
			return
		}

		// 从环境变量读取敏感的 API 密钥
		secretId := os.Getenv("TENCENTCLOUD_SECRET_ID")
		secretKey := os.Getenv("TENCENTCLOUD_SECRET_KEY")
		if secretId == "" || secretKey == "" {
			// 不返回明文错误
			response.JSON(c, eerrors.Wrap(fmt.Errorf("missing tencent cloud credentials"), code.InternalServerErr, "internal error"), nil)
			return
		}

		credential := common.NewCredential(
			secretId,
			secretKey,
		)
		cpf := profile.NewClientProfile()
		cpf.HttpProfile.Endpoint = "tiia.tencentcloudapi.com"
		client, _ := tiia.NewClient(credential, "ap-guangzhou", cpf)

		request := tiia.NewAssessQualityRequest()
		request.ImageBase64 = common.StringPtr(param.Base64Image)

		resp, err := client.AssessQuality(request)
		if _, ok := err.(*errors.TencentCloudSDKError); ok {
			// 不打印敏感错误内容到日志
			response.JSON(c, eerrors.Wrap(fmt.Errorf("tencent cloud sdk error"), code.ExternalServiceErr, "external service error"), nil)
			return
		}
		if err != nil {
			// 封装并隐藏详细错误
			response.JSON(c, eerrors.Wrap(fmt.Errorf("service error"), code.ExternalServiceErr, "external service error"), nil)
			return
		}
		c.String(http.StatusOK, resp.ToJsonString())
	}
}

// isLocalIP 检查请求的ip是否是本地ip
func isLocalIP(host string) bool {
	ip, _, err := net.SplitHostPort(host)
	if err != nil {
		return false
	}
	allowIps := []string{"localhost", "127.0.0.1"}
	for _, item := range allowIps {
		if ip == item {
			return true
		}
	}
	return false
}
