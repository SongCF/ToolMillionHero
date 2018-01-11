package ocr

import (
	"encoding/base64"
	"fmt"
	"github.com/SongCF/ToolMillionHero/baidu/auth"
	"github.com/SongCF/ToolMillionHero/utils"
	"github.com/kataras/go-errors"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"
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

func GetImageText(filename string) ([]string, error) {
	imgBytes, err := loadClipImageBytes(filename)
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
	//log.Println("ack:", ack)
	if ack.ErrorCode != 0 {
		return nil, errors.New(ack.ErrorMsg)
	}
	var l []string
	for i := range ack.WordsResult {
		l = append(l, ack.WordsResult[i].Words)
	}
	return l, nil
}

func loadImageBytes(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func loadClipImageBytes(filename string) ([]byte, error) {
	newFilename := "screenshot_clip.png"
	fIn, err := os.Open(filename)
	defer fIn.Close()
	if err != nil {
		panic(err)
	}
	fOut, _ := os.Create(newFilename)
	defer fOut.Close()

	err = clip(fIn, fOut)
	if err != nil {
		panic(err)
	}
	return ioutil.ReadFile(newFilename)
}

func clip(in io.Reader, out io.Writer) error {
	src, fm, err := image.Decode(in)
	if err != nil {
		return err
	}
	wh := src.Bounds().Max

	// 截图量的题目范围
	const w, h, board, down, top = 1080, 1920, 50, 700, 280
	realW, realH := wh.X, wh.Y
	xScale, yScale := realW/w, realH/h
	x0, y0 := board*xScale, top*yScale
	x1, y1 := realW-x0, realH-down*yScale

	switch fm {
	case "png":
		switch src.(type) {
		case *image.NRGBA:
			img := src.(*image.NRGBA)
			subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.NRGBA)
			return png.Encode(out, subImg)
		case *image.RGBA:
			img := src.(*image.RGBA)
			subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.RGBA)
			return png.Encode(out, subImg)
		default:
			return errors.New("error image type")
		}
	default:
		return errors.New("not support image format")
	}
}
