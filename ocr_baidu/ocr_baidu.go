package ocr_baidu

import (
	"github.com/songcf/ToolMillionHero/utils"
	"fmt"
	"github.com/songcf/ToolMillionHero/auth_baidu"
	"log"
	"encoding/json"
)

//网络图片文字识别
var url_image2text = "https://aip.baidubce.com/rest/2.0/ocr/v1/webimage"

type Image2TextReq struct {
	Image string `json:"image"`
}

type Image2TextAck struct {
	LogID int64 `json:"log_id"` //唯一的log id，用于问题定位
	WordsResult []string `json:"words_result"` //识别结果数组
	WordsResultNum int32 `json:"words_result_num"` //识别结果数，表示words_result的元素个数
	Words string `json:"+words"` //识别结果字符串
	Probability map[string]interface{} `json:"probability"` //识别结果中每一行的置信度值，包含average：行置信度平均值，variance：行置信度方差，min：行置信度最小值
}

func GetImageText(fileName string) {

	//
	url := fmt.Sprintf("%s?access_token=%s", url_image2text, auth_baidu.AccessToken)
	headers := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	req := &Image2TextReq{}
	ack := &Image2TextAck{}
	body, err := json.Marshal(req)
	if err != nil {
		log.Println("json marshal failed:", err)
		return
	}
	err = utils.DoHttpWithParse("POST", url, body, headers, ack)
	if err != nil {
		log.Println("DoHttpWithParse failed")
	}
	log.Println("ack:", ack)
}