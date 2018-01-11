package search

import (
	"fmt"
	"github.com/SongCF/ToolMillionHero/utils"
	"net/url"
	"regexp"
	"strings"
)

type AN struct {
	A string
	N int
}

func GetAnswerWeight(q string, l []AN) ([]AN, error) {
	urlList, err := UrlListByKey(q, 1)
	if err != nil {
		return nil, err
	}
	for _, v := range urlList {
		data, _ := getHTML(v)
		str := string(data)
		for idx := range l {
			l[idx].N += strings.Count(str, l[idx].A)
		}
	}
	return l, nil
}

func UrlListByKey(key string, pageNum int) ([]string, error) {
	key = url.QueryEscape(key)
	urlList := []string{}
	for i := 0; i < pageNum; i++ {
		////https://github.com/demonchang/baidu_search_url_list
		uri := fmt.Sprintf("http://www.baidu.com/s?ie=utf-8&mod=1&isbd=1&isid=bdbbaf61002a79be&wd=%s&pn=%d&oq=%s&tn=baiduhome_pg&ie=utf-8&usm=2&rsv_idx=2", key, i*10, key)
		body, err := utils.DoHttpGet(uri)
		if err != nil {
			return urlList, err
		}
		reg := regexp.MustCompile(`<h3[\s\S]*?href="(.*?)"[\s\S]*?<\/h3>`)
		hrefs := reg.FindAllStringSubmatch(string(body), -1)
		for _, v := range hrefs {
			urlList = append(urlList, v[1])
		}
	}
	return urlList, nil
}

func getHTML(url string) ([]byte, error) {
	return utils.DoHttpGet(url)
}
