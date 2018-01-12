package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func DoHttpWithParse(method, url string, body []byte, headerMap map[string]string, v interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		log.Println("NewRequest failed", "err", err, "url", url, "method", method, "body", string(body))
		return err
	}
	for k, v := range headerMap {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Println("client.Do failed", "err", err, "url", url, "method", method, "body", string(body), "header", headerMap)
		return err
	}
	buf, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Println("outil.ReadAll failed", "err", err, "url", url, "method", method, "body", string(body), "header", headerMap)
		return err2
	}
	//log.Println("body------->", string(buf))
	err = json.Unmarshal(buf, v)
	if err != nil {
		log.Println("json.Unmarshal failed", "err", err, "url", url, "method", method, "body", string(body), "header", headerMap, "buf", string(buf))
		return err
	}
	return nil
}

func DoHttpPostObjFormWithParse(url string, r string, v interface{}) error {
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(r))
	if err != nil {
		log.Println("http req failed", "url", url)
		log.Println("detail", "post_data", r)
		log.Println(err.Error())
		return err
	}
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Println("read resp.Body failed", "url", url)
		log.Println(err.Error())
		resp.Body.Close()
		return err2
	}
	err = resp.Body.Close()
	if err != nil {
		log.Println("read resp.Body failed", "url", url)
		log.Println(err.Error())
		return err
	}
	//log.Println("body------->", string(body))
	err = json.Unmarshal(body, v)
	if err != nil {
		log.Println("json parse failed", "info", err.Error(), "body", string(body))
		return err
	}
	return nil
}

func DoHttpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("DoHttpGet http.Get failed:", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	buf, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Println("DoHttpGet ioutil.ReadAll failed:", err2.Error())
		return nil, err2
	}
	return buf, nil
}
