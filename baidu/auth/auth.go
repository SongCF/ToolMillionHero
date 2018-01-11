package auth

import (
	"fmt"
	"github.com/SongCF/ToolMillionHero/baidu"
	"github.com/SongCF/ToolMillionHero/utils"
	"github.com/kataras/go-errors"
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
	//TODO del
	return "24.df999ff7107bdedd799ba258c752962c.2592000.1518233975.282335-10671124"
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
