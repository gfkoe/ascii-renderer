package ascii

import (
	"fmt"
	"image/color"
	//"reflect"
)

var Pixels = []byte(" .,:;i1tfLCG08@")

type CharPixel struct {
	Char       byte
	R, G, B, A uint8
}

func NewPixelConverter() PixelConverter {
	return PixelAsciiConverter{}
}

type PixelConverter interface {
	ConvertToAscii(pixel color.Color) string
	ConvertToPixelAscii(pixel color.Color) CharPixel
}

type PixelAsciiConverter struct{}

func (converter PixelAsciiConverter) ConvertToPixelAscii(pixel color.Color) CharPixel {
	r, g, b, a := pixel.RGBA()
	r8 := uint8(r >> 8)
	g8 := uint8(g >> 8)
	b8 := uint8(b >> 8)
	a8 := uint8(a >> 8)

	index := int((r + g + b) / 3 * 10 / 255)
	if index >= len(Pixels) {
		index = len(Pixels) - 1
	}
	rawPixel := Pixels[index]

	return CharPixel{
		Char: rawPixel,
		R:    r8,
		G:    g8,
		B:    b8,
		A:    a8,
	}
}

func (converter PixelAsciiConverter) ConvertToAscii(pixel color.Color) string {
	pixelVal := converter.ConvertToPixelAscii(pixel)
	rawChar, r, g, b := pixelVal.Char, pixelVal.R, pixelVal.G, pixelVal.B

	return converter.colorToPixel(r, g, b, rawChar)
}

func (converter PixelAsciiConverter) colorToPixel(r, g, b uint8, rawChar byte) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", r, g, b, string([]byte{rawChar}))
}
