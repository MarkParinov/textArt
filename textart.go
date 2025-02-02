package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, err := png.Decode(f)
	return image, err
}

func drawInText(src image.Image, width int, height int) {
	for x := range width {
		for y := range height {
			pixel := src.At(y, x)
			r, g, b, a := pixel.RGBA()

			r = r / 256
			g = g / 256
			b = b / 256
			a = a / 256

			grayscale := 0.299*float32(r) + 0.587*float32(g) + 0.114*float32(b)

			if grayscale == 0 {
				fmt.Print(" ")
			} else if grayscale < 64 {
				fmt.Print(".")
			} else if grayscale < 128 {
				fmt.Print(";")
			} else if grayscale < 192 {
				fmt.Print("/")
			} else {
				fmt.Print("@")
			}
		}
		fmt.Println()
	}
}

func main() {
	args := os.Args[1:]
	if len(args) > 1 || len(args) == 0 {
		fmt.Println("Invalid arguments. Please, provide a sigle argument that specifies path to a file.")
		return
	}
	img, err := getImageFromFilePath(args[0])
	if err != nil {
		fmt.Println("Invalid file. Please provide a path to a valid .png file.")
		return
	}
	drawInText(img, img.Bounds().Max.X, img.Bounds().Max.Y)
}
