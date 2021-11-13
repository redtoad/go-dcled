package main

import (
	"errors"
	"image"
	"image/color"
	"math"
)

var (
	On  = color.RGBA{0xff, 0x00, 0x00, 0xff}
	Off = color.RGBA{0x00, 0x00, 0x00, 0xff}
)

// NewCanvas returns a new Canvas image with the given bounds.
func NewCanvas(r image.Rectangle) *Canvas {
	bufferLength, err := pixelBufferLength(r)
	if err != nil {
		panic(err)
	}
	return &Canvas{
		Pix:    make([]uint8, bufferLength),
		Stride: 1 * r.Dx(), // in bits!
		Rect:   r,
	}
}

// pixelBufferLength calculates length of buffer in bytes. Each byte
// contains 8 pixels (=bits). If sum of bits is not a multiple of 8,
// we'll have trailing zero bits which can be ignored.
func pixelBufferLength(r image.Rectangle) (int, error) {
	if r.Empty() {
		return 0, errors.New("new canvas rectangle has zero dimension")
	}
	bits := r.Dy() * r.Dx()
	bytes := int(math.Ceil((float64(bits) / 8)))
	return bytes, nil
}

// Canvas is an in-memory image whose At method returns either
// black (off) or red (on).
type Canvas struct {
	// Pix holds the image's pixels
	Pix []uint8
	// Stride is the Pix stride (in bits!) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

// MonochromeModel can convert any Color to colours On or Off.
// The conversion is lossy!
type MonochromeModel struct{}

// Convert turns any colors which aren't black into On.
func (m MonochromeModel) Convert(c color.Color) color.Color {
	r, g, b, _ := c.RGBA()
	if r+b+g > 0 {
		return On
	}
	return Off
}

// ColorModel returns the Image's color model.
func (c *Canvas) ColorModel() color.Model {
	return MonochromeModel{}
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
// Bounds implements the Image interface.
func (c *Canvas) Bounds() image.Rectangle {
	return c.Rect
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (c Canvas) At(x, y int) color.Color {
	pixel := x + y*c.Stride
	idx := pixel / 8
	bit := pixel % 8
	bitmask := uint8(1 << bit)
	if c.Pix[idx]&bitmask == 0x00 {
		return Off
	}
	return On
}

func (c Canvas) Set(x, y int, color color.Color) {
	pixel := x + y*c.Stride
	idx := pixel / 8
	bit := pixel % 8
	bitmask := uint8(1 << bit)
	if color == On {
		c.Pix[idx] |= bitmask & 0xff
	} else if color == Off {
		c.Pix[idx] |= bitmask ^ 0x00
	}
}
