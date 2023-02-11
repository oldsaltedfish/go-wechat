package pub

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type AccessToken struct {
	ErrorMsg
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	ExpireTime  int64  `json:"expire_time"`
}

const GetAccessTokenAPi = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=${APPID}&secret=${APPSECRET}"

func getAccessToken(appid string, secret string) (*AccessToken, error) {
	url := strings.Replace(GetAccessTokenAPi, "${APPID}", appid, 1)
	url = strings.Replace(url, "${APPSECRET}", secret, 1)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	accessToken := new(AccessToken)
	err = json.Unmarshal(body, accessToken)
	if err != nil {
		return nil, err
	}
	if accessToken.ErrCode > 0 {
		return nil, errors.New(fmt.Sprintf("code: %v, msg: %v", accessToken.ErrCode, accessToken.ErrMsg))
	}
	accessToken.ExpireTime = time.Now().Add(time.Duration(accessToken.ExpiresIn)*time.Second - 10*time.Minute).Unix()
	return accessToken, nil
}
