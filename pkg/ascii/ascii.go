package ascii

import (
	"image/color"
	"reflect"
)

type Char struct {
	Char       byte
	R, G, B, A uint8
}

type PixelConverter interface {
	ConvertToAscii(pixel color.Color) string
	ConvertToPixelAscii(pixel color.Color) Char
}

type PixelAsciiConverter struct{}

func (converter PixelAsciiConverter) ConvertToPixelAscii(pixel color.Color) Char {
	r := uint8(reflect.ValueOf(pixel).FieldByName("R").Uint())
	g := uint8(reflect.ValueOf(pixel).FieldByName("G").Uint())
	b := uint8(reflect.ValueOf(pixel).FieldByName("B").Uint())
	a := uint8(reflect.ValueOf(pixel).FieldByName("A").Uint())
	return Char{
		Char: ' ',
		R:    uint8(r),
		G:    uint8(g),
		B:    uint8(b),
		A:    uint8(a),
	}
}

func (converter PixelAsciiConverter) ConvertToAscii(pixel color.Color) string {
	pixelVal := converter.ConvertToPixelAscii(pixel)
	char := pixelVal.Char
	return string([]byte{char})
}
