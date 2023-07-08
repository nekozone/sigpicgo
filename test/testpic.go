package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	_ "image/png"
	"math/rand"
	"os"
	"time"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

var picfile = "assent/bd3.png"

// var ximg image.Image
var font *truetype.Font
var texts = []string{"忽有狂徒夜磨刀，帝星飘摇荧惑高。", "IP地址: {}  {}", "https://neko.red", "浏览器: {} OS: {} 设备: {}", "No. 112358", "@dogcraft neko.red"}
var picsrclist = []string{"assent/bd.png", "assent/bd2.png", "assent/bd3.png"}
var imglist []image.Image
var piclistnum = 0

// var rawpic []byte

func pinit() {
	fmt.Println("init")
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(picsrclist); i++ {
		input, err := os.Open(picsrclist[i])
		if err != nil {
			panic(err)
		}
		defer input.Close()
		img, _, err := image.Decode(input)
		if err != nil {
			panic(err)
		}
		imglist = append(imglist, img)
	}
	piclistnum = len(imglist)
	fontfile, err := os.ReadFile("assent/LXGWWenKaiUI-Regular.ttf")
	if err != nil {
		panic(err)
	}
	// defer fontfile.Close()
	font, err = freetype.ParseFont(fontfile)
	if err != nil {
		panic(err)
	}

}

func pmain() {
	fmt.Println("Hello World!")
	// fmt.Println(ximg.Bounds())
	// textpic("dogcraft123、你好sssssss")
	textpic("你好")
	textpic("abcdog")
	textpic("dogcraft")

}

func randcolor1() *image.Uniform {
	return image.NewUniform(color.RGBA{uint8(rand.Intn(144)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255})
}

func textpic(text string) {
	// var newraw []byte
	// newraw := bytes.Clone(rawpic)
	// newimg, _, err := image.Decode(bytes.NewReader(newraw))
	// // var dd draw.Image
	// draw.Draw(dd, newimg.Bounds(), newimg, image.ZP, draw.Src)
	ximg := imglist[rand.Intn(piclistnum)]
	newimg := image.NewRGBA(ximg.Bounds())
	// newimg := image.NewRGBA(image.Rect(0, 0, 1020, 240))
	draw.Draw(newimg, ximg.Bounds(), ximg, image.Point{}, draw.Src)
	c := freetype.NewContext()
	c.SetFont(font)
	c.SetFontSize(22)
	c.SetDst(newimg)
	c.SetClip(newimg.Bounds())
	// draw.text((10, 30-10), dog_text_1, font=font2, fill=(255,0,0))
	// draw.text((10, 70-10), dog_text_2, font=font2, fill=(0,0,0))
	// draw.text((10, 110-8), dog_text_3, font=font1, fill=rndColor2())
	// draw.text((10, 150-10), dog_text_4, font=font2, fill=rndColor2())
	// draw.text((10, 180-5), dog_text_5, font=font2, fill=rndColor2())
	// draw.text((10, 210),dog_text_6, font=font2, fill=rndColor2())
	c.SetSrc(image.NewUniform(color.RGBA{255, 0, 0, 255}))
	c.DrawString(texts[0], freetype.Pt(10, 20+20))
	c.SetSrc(image.NewUniform(color.RGBA{0, 0, 0, 255}))
	c.DrawString(texts[1], freetype.Pt(10, 60+20))
	c.SetFontSize(16)
	c.SetSrc(randcolor1())
	c.DrawString(texts[2], freetype.Pt(10, 102+20))
	c.SetSrc(randcolor1())
	c.SetFontSize(22)
	c.DrawString(texts[3], freetype.Pt(10, 140+20))
	c.SetSrc(randcolor1())
	c.DrawString(texts[4], freetype.Pt(10, 175+20))
	c.SetSrc(randcolor1())
	c.DrawString(texts[5], freetype.Pt(10, 210+20))

	// fcolor := image.NewUniform(randcolor1())
	// c.SetSrc(fcolor)
	// _, err := c.DrawString(text, freetype.Pt(50, 50))
	// if err != nil {
	// 	panic(err)
	// }
	// c.SetFontSize(16)
	// _, err = c.DrawString(fmt.Sprintf("%s", time.Now().UTC()), freetype.Pt(70, 100))
	// if err != nil {
	// 	panic(err)
	// }
	scl := time.Now().UnixMicro()
	output, err := os.Create(fmt.Sprintf("assent/%d.jpg", scl))
	if err != nil {
		panic(err)
	}
	fmt.Println(newimg.Bounds())
	defer output.Close()
	b := bufio.NewWriter(output)
	err = jpeg.Encode(b, newimg, nil)
	if err != nil {
		panic(err)
	}
}
