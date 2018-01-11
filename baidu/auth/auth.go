package auth

import (
	"fmt"
	"github.com/SongCF/ToolMillionHero/baidu"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var expireTime = int64(86400) * 10 //默认10天token超市，官网提供的时30天
var Token = &struct {
	AccessToken string
	Time        int64
}{}

func AccessToken() string {
	now := time.Now().Unix()
	if now-Token.Time > expireTime {
		initToken()
	}
	return Token.AccessToken
}

func initToken() {
	token, err := getAccessToken()
	if err != nil {
		log.Println("Error: getAccessToken failed")
		panic(err)
	}
	Token.AccessToken = token
	Token.Time = time.Now().Unix()
}

func getAccessToken() (string, error) {
	url := fmt.Sprintf("https://aip.baidubce.com/oauth/2.0/token?grant_type=%s&client_id=%s&client_secret=%s",
		"client_credentiale", baidu.AppKey, baidu.SecretKey)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	buf, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return "", err2
	}
	return string(buf), nil
}
