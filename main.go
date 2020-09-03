package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"image/jpeg"
	"io/ioutil"
	"math"
	"os"
	"time"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
	"golang.org/x/image/colornames"
)

var dtf = "2006-01-02 15:04:05.000"
var font = loadFont()
var face50 = truetype.NewFace(font, &truetype.Options{Size: 50})
var face26 = truetype.NewFace(font, &truetype.Options{Size: 26})
var face29 = truetype.NewFace(font, &truetype.Options{Size: 29})
var backgroundImg, _ = gg.LoadImage("image_background.png")

func main() {
	//var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 10; i++ {
		drawContext := draw()
		//wg.Add(1)
		//go func() {
		file, err := os.OpenFile("out_"+uuid.NewV4().String()+".png", os.O_CREATE, os.ModeAppend)
		handleError(err)
		jpeg.Encode(file, drawContext.Image(), nil)
		fmt.Println(time.Now().Format(dtf), "save img")
		//	defer wg.Done()
		//}()
	}
	//wg.Wait()
	end := time.Now()
	fmt.Println(time.Now().Format(dtf), end.Sub(start))
}

func draw() *gg.Context {
	fmt.Println(time.Now().Format(dtf), "demo2")

	const width = 818
	const height = 1057
	dc := gg.NewContext(width, height)
	dc.SetRGBA(0, 0, 0, 0.1)
	fmt.Println(time.Now().Format(dtf), "NewContext")

	dc.DrawImage(backgroundImg, 0, 0)
	fmt.Println(time.Now().Format(dtf), "backgroundImg")

	dc.SetColor(colornames.Red)
	dc.SetFontFace(face50)
	dc.DrawString("文善", 312, float64(246+face50.Metrics().Ascent.Ceil()))
	fmt.Println(time.Now().Format(dtf), "文善")

	dc.SetRGB(51, 51, 51)
	dc.SetFontFace(face26)
	dc.DrawString("客户经理", 312, float64(313+face26.Metrics().Ascent.Ceil()))
	fmt.Println(time.Now().Format(dtf), "客户经理")
	dc.Fill()

	dc.SetRGB(51, 51, 51)
	dc.SetFontFace(face29)
	dc.DrawString("杭州君方科技有限公司", 364, float64(393+face29.Metrics().Ascent.Ceil()))
	fmt.Println(time.Now().Format(dtf), "杭州君方科技有限公司")

	dc.SetRGB(51, 51, 51)
	dc.SetFontFace(face29)
	dc.DrawString("130 9998 7861", 364, float64(446+face29.Metrics().Ascent.Ceil()))
	fmt.Println(time.Now().Format(dtf), "130 9998 7861")

	text := "专业生产销售环保涤纶彩色化纤、钻石丝、炫彩丝、纺尼龙丝、彩色水晶丝、彩色毛毯丝、彩色低弹高F彩色、彩色复合系列、阳涤复合原料、热溶及特种丝各类定纺。"
	dc.SetRGB(51, 51, 51)
	dc.SetFontFace(face26)
	limit := int(math.Ceil(531 * 1.0 / 26))
	index := 0
	textArr := []rune(text)
	lineHeight := 0
	for {
		from := index
		to := limit + index

		if from > len(textArr) {
			break
		}

		if to > len(textArr) {
			to = len(textArr)
		}

		fmt.Println(time.Now().Format(dtf), "from ", from, " ", "to ", to, " ", "limit ", limit, " ", string(textArr[from:to]))

		arr := textArr[from:to]

		dc.DrawString(string(textArr[from:to]), 149, float64(621+lineHeight+face26.Metrics().Ascent.Ceil()))

		if len(arr) == 0 {
			break
		}
		index += limit
		lineHeight += 35
	}
	fmt.Println(time.Now().Format(dtf), "text")

	dc.SetRGB(0, 0, 0)
	dc.SetFontFace(face29)
	dc.DrawString("扫码关注", 460, float64(865+face29.Metrics().Ascent.Ceil()))
	fmt.Println(time.Now().Format(dtf), "扫码关注")

	avatarImg, err := gg.LoadImage("image_avatar.png")
	avatarImg = resize.Resize(152, 152, avatarImg, resize.Lanczos3)
	handleError(err)
	dc.DrawRoundedRectangle(124, 237, 152, 152, 76)
	dc.Clip()
	dc.DrawImage(avatarImg, 124, 237)
	fmt.Println(time.Now().Format(dtf), "image_avatar")

	dc.Fill()
	return dc
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func loadFont() *truetype.Font {
	fontBytes, err := ioutil.ReadFile("PINGFANG.ttf")
	font, err := truetype.Parse(fontBytes)
	handleError(err)
	return font
}
