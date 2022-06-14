package ping

import (
	eerrors "BloodPressure/pkg/errors"
	"BloodPressure/pkg/errors/code"
	"BloodPressure/pkg/response"
	"fmt"
	"net"
	"net/http"

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

		credential := common.NewCredential(
			"AKIDJcnocF1nQxJXZO0NOGwTt7D9uUt54vZz",
			"GnBRVSL3f9TfhbbOyUYS8c7axFRZcaUK",
		)
		cpf := profile.NewClientProfile()
		cpf.HttpProfile.Endpoint = "tiia.tencentcloudapi.com"
		client, _ := tiia.NewClient(credential, "ap-guangzhou", cpf)

		request := tiia.NewAssessQualityRequest()

		request.ImageBase64 = common.StringPtr(param.Base64Image)

		response, err := client.AssessQuality(request)
		if _, ok := err.(*errors.TencentCloudSDKError); ok {
			fmt.Printf("An API error has returned: %s", err)
			return
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s", response.ToJsonString())
		c.String(http.StatusOK, response.ToJsonString())
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
