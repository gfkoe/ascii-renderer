package render

import (
	"ascii-renderer/pkg/ascii"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func NewImageConverter() *ImageConverter {
	return &ImageConverter{
		pixelConverter: ascii.NewPixelConverter(),
	}
}

type ImageConverter struct {
	pixelConverter ascii.PixelConverter
}

func (converter *ImageConverter) RenderImage(path string) (string, error) {
	img, err := OpenImage(path)
	if err != nil {
		log.Fatal(err)
	}
	sz := img.Bounds()
	w := sz.Max.X
	h := sz.Max.Y
	asciis := make([][]ascii.CharPixel, 0, h)

	for i := 0; i < int(h); i++ {
		line := make([]ascii.CharPixel, 0, w)
		for j := 0; j < int(w); j++ {
			pixel := color.NRGBAModel.Convert(img.At(j, i))
			pixelAscii := converter.pixelConverter.ConvertToPixelAscii(pixel)
			line = append(line, pixelAscii)
		}
		asciis = append(asciis, line)
	}

	// Create a string representation of the ASCII art
	result := ""
	for _, line := range asciis {
		for _, pixel := range line {
			result += string(pixel.Char)
		}
		result += "\n" // Add a new line after each row
	}

	return result, nil
}

func OpenImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}
