package openid

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/thedevsaddam/gojsonq/v2"
)

const (
	code2sessionURL = ""
	// appID最好不要直接暴露，这里以后一定要改 - 2022.5.25日留
	appID     = ""
	appSecret = ""
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
