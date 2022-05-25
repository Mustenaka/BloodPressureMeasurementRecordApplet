package openid

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/thedevsaddam/gojsonq/v2"
)

const (
	code2sessionURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	appID           = "wxa2164c390f287f1c"
	appSecret       = "8493e5f7166702c75143553ea1eaf502"
)

// GetOpenID 通过用户code获取openid
func GetOpenidByCode(code string) (string, error) {
	url := fmt.Sprintf(code2sessionURL, appID, appSecret, code)
	resp, err := http.Get(url)
	if err != nil {
		return "", errors.New("get openid failed")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("read openid failed")
	}
	fmt.Println(string(body))
	json := gojsonq.New().FromString(string(body)).Find("openid")
	openId := json.(string)
	return openId, nil
}
