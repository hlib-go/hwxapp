package hwxapp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html

// 登录凭证校验。通过 wx.login 接口获得临时登录凭证 code 后传到开发者服务器调用此接口完成登录流程。
func Jscode2session(cfg *Config, code string) (r *Jscode2sessionResult, err error) {
	resp, err := http.Get(fmt.Sprintf("GET https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", cfg.Appid, cfg.Secret, code))
	if err != nil {
		return
	}
	rbytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(rbytes, &r)
	if err != nil {
		return
	}
	return
}

type Jscode2sessionResult struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	Errcode    string `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}
