package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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
	err = json.Unmarshal(buf, v)
	if err != nil {
		log.Println("json.Unmarshal failed", "err", err, "url", url, "method", method, "body", string(body), "header", headerMap, "buf", string(buf))
		return err
	}
	return nil
}
