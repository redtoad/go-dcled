package dcled

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

// IsOn will return true if Color c represents on.
func IsOn(c color.Color) bool {
	r, g, b, _ := c.RGBA()
	onR, onG, onB, _ := On.RGBA()
	return r == onR && g == onG && b == onB
}

// IsOn will return true if Color c represents off.
func IsOff(c color.Color) bool {
	r, g, b, _ := c.RGBA()
	offR, offG, offB, _ := Off.RGBA()
	return r == offR && g == offG && b == offB
}

// MonochromeModel can convert any Color to colours On or Off.
// The conversion is lossy!
var MonochromeModel = color.ModelFunc(monochromeModel)

func monochromeModel(c color.Color) color.Color {
	r, g, b, _ := c.RGBA()
	if r+b+g > 0 {
		return On
	}
	return Off
}

// Canvas is an in-memory image whose At method returns either
// black (off) or red (on).
type Canvas interface {
	Bounds() image.Rectangle
	At(x, y int) color.Color
	SubImage(r image.Rectangle) image.Image
}

func canvasToGrid(img image.Image) [][]int {
	grid := make([][]int, 7)
	xOff := img.Bounds().Min.X
	yOff := img.Bounds().Min.Y
	for y := 0; y < 7; y++ {
		grid[y] = make([]int, 21)
		for x := 0; x < 21; x++ {
			if IsOn(img.At(x+xOff, y+yOff)) {
				grid[y][x] = 1
			} else {
				grid[y][x] = 0
			}
		}
	}
	return grid
}

// DisplayCanvas displays an image.Image on the Dream Cheeky
// Message board. Pixels of color #ff0000 (On) will turn the
// respective leds on, other color will turn them off.
func DisplayCanvas(img image.Image, device *hid.Device) error {
	grid := canvasToGrid(img)
	return DisplayGrid(grid, device)
}
