package main

import (
	"ascii-renderer/pkg/render"
	"flag"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

var imagePath string

func init() {
	flag.StringVar(&imagePath, "f", "", "Path to the image file")
}

func main() {
	flag.Parse()
	converter := render.NewImageConverter()

	output, err := converter.RenderImage(imagePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the ASCII art
	fmt.Print(output)
}
