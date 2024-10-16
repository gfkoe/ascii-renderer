package ascii

import (
	"image/color"
	"reflect"
)

type Char struct {
	Char       rune
	R, G, B, A uint8
}

func ConvertToAscii(pixel color.Color) Char {
	return Char{
		Char: ' ',
		R:    uint8(reflect.ValueOf(pixel).FieldByName("R").Uint()),
		G:    uint8(reflect.ValueOf(pixel).FieldByName("G").Uint()),
		B:    uint8(reflect.ValueOf(pixel).FieldByName("B").Uint()),
		A:    uint8(reflect.ValueOf(pixel).FieldByName("A").Uint()),
	}
}
