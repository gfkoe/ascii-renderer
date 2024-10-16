package ascii

import (
	"image/color"
	"reflect"
	"testing"
)

func TestCConvertToAscii(t *testing.T) {
	// Test cases
	tests := []struct {
		name   string
		pixel  color.Color
		expect Char
	}{
		{
			name:  "Black",
			pixel: color.RGBA{0, 0, 0, 255},
			expect: Char{
				Char: ' ',
				R:    0,
				G:    0,
				B:    0,
				A:    255,
			},
		},
		{
			name:  "White",
			pixel: color.RGBA{255, 255, 255, 255},
			expect: Char{
				Char: ' ',
				R:    255,
				G:    255,
				B:    255,
				A:    255,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Execute the function
			got := ConvertToAscii(tt.pixel)

			// Compare the result
			if !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("unexpected result: got %v, expect %v", got, tt.expect)
			}
		})
	}
}
