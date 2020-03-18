package helper

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image implements image.Image interface
type Image struct {
	Width  int
	Height int
}

// ColorModel is defined as RGBAModel
func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns rectangle
func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.Width, i.Height)
}

// At method returns color based on location
func (i *Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x) * uint8(y), uint8(x) * uint8(y), 255, 255}
}

// ShowImage prints image
func ShowImage() {
	m := &Image{Width: 255, Height: 255}
	pic.ShowImage(m)
}
