package auth

import (
	"errors"
	"fmt"
	"github.com/SongCF/ToolMillionHero/baidu"
	"github.com/SongCF/ToolMillionHero/utils"
	"log"
	"time"
)

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	Error       string `json:"error"`
	ErrDesc     string `json:"error_description"`
	TS          int64
}

var _token Token

func AccessToken() string {
	return "24.cf28fda5adc3b59e7e4c38f3d1dd73d9.2592000.1518316427.282335-10671124"
	now := time.Now().Unix()
	if now-_token.TS >= _token.ExpiresIn {
		initToken()
	}
	return _token.AccessToken
}

func initToken() {
	t, err := getAccessToken()
	if err != nil {
		log.Println("Error: getAccessToken failed")
		panic(err)
	}
	log.Println(*t)
	_token = *t
	_token.TS = time.Now().Unix()
}

func getAccessToken() (*Token, error) {
	url := fmt.Sprintf("https://aip.baidubce.com/oauth/2.0/token?grant_type=%s&client_id=%s&client_secret=%s",
		"client_credentials", baidu.AppKey, baidu.SecretKey)
	ack := &Token{}
	err := utils.DoHttpWithParse("POST", url, nil, nil, ack)
	if ack.Error != "" {
		return ack, errors.New("getAccessToken failed: [error]" + ack.Error + " [desc]" + ack.ErrDesc)
	}
	return ack, err
}
