package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {

	width := 512
	height := 512

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	red := color.RGBA{255, 0, 0, 0xff}
	green := color.RGBA{0, 255, 0, 0xff}
	blue := color.RGBA{0, 0, 255, 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2: // upper left quadrant
				img.Set(x, y, red)
			case x >= width/2 && y >= height/2: // lower right quadrant
				img.Set(x, y, blue)
			default:
				img.Set(x, y, green)
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create("output/drawing_01.png")
	png.Encode(f, img)
}
