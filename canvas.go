package main

import (
	"image"
	"image/color"

	"github.com/karalabe/hid"
)

// Colors used when displaying Canvas as image.
var (
	On  = color.RGBA{0xff, 0x00, 0x00, 0xff} // red
	Off = color.RGBA{0x00, 0x00, 0x00, 0xff} // black
)

// MonochromeModel can convert any Color to colours On or Off.
// The conversion is lossy!
var MonochromeModel = color.ModelFunc(monochromeModel)

// Convert turns any colors which aren't black into On.
func monochromeModel(c color.Color) color.Color {
	r, g, b, _ := c.RGBA()
	if r+b+g > 0 {
		return On
	}
	return Off
}

// NewCanvas returns a new Canvas image with the given bounds.
func NewCanvas(r image.Rectangle) *Canvas {
	return &Canvas{image.NewNRGBA(r)}
}

// Canvas is an in-memory image whose At method returns either
// black (off) or red (on).
type Canvas struct {
	*image.NRGBA
}

// ColorModel returns the Image's color model.
func (c *Canvas) ColorModel() color.Model {
	return MonochromeModel
}

func match(col1, col2 color.Color) bool {
	r1, g1, b1, a1 := col1.RGBA()
	r2, g2, b2, a2 := col2.RGBA()
	return r1 == r2 && g1 == g2 && b1 == b2 && a1 == a2
}

func canvasToGrid(canvas image.Image) [][]int {
	grid := make([][]int, 7)
	for y := 0; y < 7; y++ {
		grid[y] = make([]int, 21)
		for x := 0; x < 21; x++ {
			if match(canvas.At(x+1, y+1), On) {
				grid[y][x] = 1
			} else {
				grid[y][x] = 0
			}
		}
	}
	return grid
}

func DisplayCanvas(canvas *Canvas, device *hid.Device) error {
	grid := canvasToGrid(canvas)
	return DisplayGrid(grid, device)
}
