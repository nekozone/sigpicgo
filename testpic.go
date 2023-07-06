package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

var picfile = "assent/bd2.png"
var ximg image.Image
var font *truetype.Font

func init() {
	fmt.Println("init")
	input, err := os.Open(picfile)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	img, _, err := image.Decode(input)
	if err != nil {
		panic(err)
	}
	ximg = img
	fontfile, err := os.ReadFile("assent/LXGWWenKaiUI-Regular.ttf")
	if err != nil {
		panic(err)
	}
	font, err = freetype.ParseFont(fontfile)
	if err != nil {
		panic(err)
	}

}

func main() {
	fmt.Println("Hello World!")
	fmt.Println(ximg.Bounds())
	textpic("你好")
	textpic("abcdog")
	textpic("dogcraft")

}

func textpic(text string) {
	newimg := image.NewRGBA(ximg.Bounds())
	draw.Draw(newimg, ximg.Bounds(), ximg, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetFont(font)
	c.SetFontSize(20)
	c.SetDst(newimg)
	c.SetClip(newimg.Bounds())
	fcolor := image.NewUniform(color.RGBA{255, 0, 0, 255})
	c.SetSrc(fcolor)
	_, err := c.DrawString(text, freetype.Pt(30, 30))
	if err != nil {
		panic(err)
	}
	output, err := os.Create(fmt.Sprintf("assent/%s.png", text))
	if err != nil {
		panic(err)
	}
	defer output.Close()
	b := bufio.NewWriter(output)
	err = png.Encode(b, newimg)
	if err != nil {
		panic(err)
	}
}
