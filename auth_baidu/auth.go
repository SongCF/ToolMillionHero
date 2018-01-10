package auth_baidu

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
)

var AccessToken string

func Init() {
	var err error
	AccessToken, err = GetAccessToken()
	if err != nil {
		log.Println("GetAccessToken failed")
		panic(err)
	}
}

func GetAccessToken() (string, error){
	url := fmt.Sprintf("https://aip.baidubce.com/oauth/2.0/token?grant_type=%s&client_id=%s&client_secret=%s",
		"client_credentiale", AppKey, SecretKey)
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
