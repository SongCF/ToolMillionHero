package main

import (
	"fmt"
	"github.com/SongCF/ToolMillionHero/baidu/ocr"
	"github.com/SongCF/ToolMillionHero/baidu/search"
	"log"
	"os/exec"
	"sort"
	"strings"
	"time"
)

func main() {
	const filename = "screenshot.png"
	//Screenshot(filename) // adb shell 在golang里用不了
	bt := time.Now()
	words, err := ocr.GetImageText(filename)
	log.Println("ts: GetImageText", time.Since(bt))
	if err != nil {
		panic(err)
	}
	bt = time.Now()
	q, aList := getQuestion(words)
	l := []search.AN{}
	for _, v := range aList {
		l = append(l, search.AN{A: v, N: 0})
	}
	log.Println("ts: getQuestion", time.Since(bt))
	bt = time.Now()
	l, err = search.GetAnswerWeight(q, l)
	log.Println("ts: GetAnswerWeight", time.Since(bt))
	if err != nil {
		log.Println("search.GetAnswerWeight failed:", err)
	}
	//sort
	bt = time.Now()
	sort.Slice(l, func(i, j int) bool {
		return l[i].N < l[j].N
	})
	log.Println("ts: sort", time.Since(bt))
	fmt.Println("Q:", q)
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

//func Screenshot(filename string) {
//	_, err := exec.Command("/system/bin/screencap", "-p", filename).Output()
//	if err != nil {
//		panic("screenshot failed")
//	}
//}

func Screenshot(filename string) {
	var str string
	var cmd *exec.Cmd
	var err error
	str = fmt.Sprintf("adb shell /system/bin/screencap -p /data/local/tmp/%s", filename)
	cmd = exec.Command(str)
	err = cmd.Run()
	if err != nil {
		panic("screenshot failed:" + err.Error())
	}
	str = fmt.Sprintf("adb pull /data/local/tmp/%s ./", filename)
	cmd = exec.Command(str)
	err = cmd.Run()
	if err != nil {
		panic("pull failed:" + err.Error())
	}
}
