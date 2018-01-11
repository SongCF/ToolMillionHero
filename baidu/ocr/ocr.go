package ocr

import (
	"encoding/base64"
	"fmt"
	"github.com/SongCF/ToolMillionHero/baidu/auth"
	"github.com/SongCF/ToolMillionHero/utils"
	"io/ioutil"
	"log"
	"net/url"
)

//网络图片文字识别
var url_image2text = "https://aip.baidubce.com/rest/2.0/ocr/v1/webimage"

type Image2TextAck struct {
	ErrorCode      int64                  `json:"error_code"`
	ErrorMsg       string                 `json:"error_msg"`
	LogID          int64                  `json:"log_id"`           //唯一的log id，用于问题定位
	WordsResult    []Words                `json:"words_result"`     //识别结果数组
	WordsResultNum int32                  `json:"words_result_num"` //识别结果数，表示words_result的元素个数
	Words          string                 `json:"+words"`           //识别结果字符串
	Probability    map[string]interface{} `json:"probability"`      //识别结果中每一行的置信度值，包含average：行置信度平均值，variance：行置信度方差，min：行置信度最小值
}

type Words struct {
	Words string `json:"words"`
}

func GetImageText(filename string) {
	imgBytes, err := loadImageBytes(filename)
	img := base64.StdEncoding.EncodeToString(imgBytes)
	log.Println("image size:", len(img))
	//
	path := fmt.Sprintf("%s?access_token=%s", url_image2text, auth.AccessToken())
	form := url.Values{}
	form.Add("image", img)
	body := form.Encode()
	ack := &Image2TextAck{}
	err = utils.DoHttpPostObjFormWithParse(path, body, ack)
	if err != nil {
		log.Println("DoHttpPostObjFormWithParse failed")
	}
	log.Println("ack:", ack)
}

func loadImageBytes(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

//func loadImage(filename string) image.Image {
//	inFile, err := os.Open(filename)
//	if err != nil {
//		panic(err)
//	}
//	src, err := png.Decode(inFile)
//	if err != nil {
//		panic(err)
//	}
//	inFile.Close()
//	return src
//}
