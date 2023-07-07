package pic

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"sigpicgo/global"

	"github.com/golang/freetype"
)

func Genpic(texts [6]string) []byte {
	ximg := global.Piclist[rand.Intn(global.Piclistnum)]
	newimg := image.NewRGBA(ximg.Bounds())
	draw.Draw(newimg, ximg.Bounds(), ximg, image.Point{}, draw.Src)
	c := freetype.NewContext()
	c.SetFont(global.Font)
	c.SetFontSize(22)
	c.SetDst(newimg)
	c.SetClip(newimg.Bounds())
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
	buf := new(bytes.Buffer)
	// defer buf.Reset()
	err := png.Encode(buf, newimg)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func randcolor1() *image.Uniform {
	return image.NewUniform(color.RGBA{uint8(rand.Intn(144)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255})
}
