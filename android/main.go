package main

import (
	"fmt"
	"github.com/SongCF/ToolMillionHero/baidu/ocr"
)

func main() {
	ocr.GetImageText("../screenshot_1.png")
	fmt.Println("ok")
}
