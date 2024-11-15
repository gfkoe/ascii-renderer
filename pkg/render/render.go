package render

import (
	"ascii-renderer/pkg/ascii"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"strings"
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

	var result strings.Builder

	for i := 0; i < int(h); i++ {
		for j := 0; j < int(w); j++ {
			pixel := color.NRGBAModel.Convert(img.At(j, i))
			pixelAscii := converter.pixelConverter.ConvertToAscii(pixel)
			result.WriteString(pixelAscii)
		}
		result.WriteString("\n")
	}

	return result.String(), nil
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
