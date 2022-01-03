package dcled

import (
	"image"
	"image/color"
	"image/draw"
	"time"
)

// Colors used when displaying Canvas as image.
var (
	On  = color.RGBA{0xff, 0x00, 0x00, 0xff} // red
	Off = color.RGBA{0x00, 0x00, 0x00, 0xff} // black
)

const (
	DeviceHeight = 7
	DeviceWidth  = 21
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

func canvasToGrid(img image.Image) [][]int {
	grid := make([][]int, DeviceHeight)
	xOff := img.Bounds().Min.X
	yOff := img.Bounds().Min.Y
	for y := 0; y < DeviceHeight; y++ {
		grid[y] = make([]int, DeviceWidth)
		for x := 0; x < DeviceWidth; x++ {
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
func DisplayCanvas(img image.Image, device Device) error {
	grid := canvasToGrid(img)
	return DisplayGrid(grid, device)
}

// NewCanvas returns a new blank image of 21x7 pixels.
func NewCanvas() draw.Image {
	img := image.NewRGBA(image.Rect(0, 0, DeviceWidth, DeviceHeight))
	draw.Draw(img, img.Bounds(), &image.Uniform{Off}, image.Point{0, 0}, draw.Src)
	return img
}

// Scroll will scroll an image across the device.
func Scroll(dev Device, img draw.Image) {
	x := 0
	dir := 1
	dst := NewCanvas()
	sr := img.Bounds()

	for {
		r := sr.Sub(image.Point{x, 0})
		draw.Draw(dst, r, img, sr.Min, draw.Src)
		_ = DisplayCanvas(dst, dev)
		time.Sleep(50 * time.Millisecond)

		x += dir
		if x >= sr.Dx() {
			x = 0
		}
	}
}

// Center will center an image on the Device.
func Center(dev Device, img draw.Image) {
	sr := img.Bounds()
	width := sr.Dx()
	height := sr.Dy()

	dx := (DeviceWidth - width) / 2
	dy := (DeviceHeight - height) / 2

	dst := NewCanvas()
	r := sr.Add(image.Point{dx, dy})
	draw.Draw(dst, r, img, sr.Min, draw.Src)
	_ = DisplayCanvas(dst, dev)
}
