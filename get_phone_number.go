package hwxapp

import (
	"encoding/json"
	"errors"
	"strings"
)

// 解密小程序返回的密文手机号  https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/getPhoneNumber.html
func GetPhoneNumber(conf *Config, sessionKey, iv, encryptedData string) (r *PhoneNumberResult, err error) {
	rbytes, err := DecryptAesCbcPKCS7(sessionKey, iv, encryptedData)
	if err != nil {
		return
	}
	if !strings.Contains(string(rbytes), conf.Appid) {
		err = errors.New("ERROR:取微信绑定手机号返回的appid与实际appid不一致")
		return
	}
	err = json.Unmarshal(rbytes, &r)
	if err != nil {
		return
	}
	return
}

type PhoneNumberResult struct {
	PhoneNumber     string `json:"phoneNumber"`
	PurePhoneNumber string `json:"purePhoneNumber"`
	CountryCode     string `json:"countryCode"`
}
