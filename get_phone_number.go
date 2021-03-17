package hwxapp

import (
	"encoding/json"
)

// 解密小程序返回的密文手机号  https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/getPhoneNumber.html
func GetPhoneNumber(appid, sessionKey, iv, encryptedData string) (r *PhoneNumberResult, err error) {
	rbytes, err := DecryptData(appid, sessionKey, encryptedData, iv)
	if err != nil {
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

func (o *PhoneNumberResult) JSON() string {
	pbs, _ := json.Marshal(o)
	return string(pbs)
}
