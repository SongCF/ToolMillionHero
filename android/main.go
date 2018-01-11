package main

import (
	"fmt"
	"github.com/SongCF/ToolMillionHero/baidu/ocr"
	"github.com/SongCF/ToolMillionHero/baidu/search"
	"log"
	"os/exec"
	"sort"
	"strings"
)

func main() {
	const filename = "screenshot.png"
	Screenshot(filename)
	words, err := ocr.GetImageText(filename)
	if err != nil {
		panic(err)
	}
	q, aList := getQuestion(words)
	l := []search.AN{}
	for _, v := range aList {
		l = append(l, search.AN{A: v, N: 0})
	}
	l, err = search.GetAnswerWeight(q, l)
	if err != nil {
		log.Println("search.GetAnswerWeight failed:", err)
	}
	//sort
	sort.Slice(l, func(i, j int) bool {
		return l[i].N < l[j].N
	})
	for _, v := range l {
		fmt.Printf("%d  %s\n", v.N, v.A)
	}
}

func getQuestion(words []string) (string, []string) {
	q := ""
	for i, v := range words {
		q += v
		if strings.Contains(v, "?") || strings.Contains(v, "？") {
			return q, words[i+1:]
		}
	}
	log.Println("getQuestion failed")
	return q, words[1:] //第一个肯定不是答案， 为了搜索q还是返回全部
}

func Screenshot(filename string) {
	_, err := exec.Command("/system/bin/screencap", "-p", filename).Output()
	if err != nil {
		panic("screenshot failed")
	}
}

//func Screenshot(filename string) {
//	cmd := fmt.Sprintf("adb shell /system/bin/screencap -p /data/local/tmp/%s", filename)
//	_, err := exec.Command("adb shell /system/bin/screencap", "-p", filename).Output()
//	if err != nil {
//		panic("screenshot failed:" + err.Error())
//	}
//	cmd = fmt.Sprintf("adb pull /data/local/tmp/%s ./", filename)
//	_, err = exec.Command(cmd).Output()
//	if err != nil {
//		panic("pull failed:" + err.Error())
//	}
//}
