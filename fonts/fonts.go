// Package fonts provides functions to write text in specific fonts
// to an image which can then be display on the LED message board.
package fonts

import (
	"image"
	"time"

	"github.com/redtoad/go-dcled"
)

//go:generate go run ./generator/gen.go *.dlf

/*
	From the comment in the original code:

	Font is 7 bytes per entry, each byte is a row.  The character bitmaps are
	like 5 bits wide, mirrored, starting at bit zero.  Why so bizzare, you
	ask?  Oh god, the horror of converting from an existing font to this...
	ImageMagick -draw text, conversion to xbm format, tcl scripts to parse the
	xbm data into this c code... anyway, it was faster than drawing a font
	myself, although not by much.  Ulitmately, this array was built
	automatically from the X11 5x7 font, and it works.  That is what matters
	for an afternoon project.  Feel free to improve it. :)
*/
type Font struct {
	Name  string
	Meta  map[string]string
	Chars [][]byte
}

// Text creates an image out of a string which can be displayed.
func Text(txt string, font Font) image.Image {

	charWidth := 5
	charHeight := 7
	gap := 0

	// We'll convert this into an array of runes,
	// otherwise the count is off for unicode character.
	characters := []rune(txt)

	width := (charWidth + gap) * len(characters)
	img := image.NewNRGBA(image.Rect(0, 0, width, charHeight))

	for pos := 0; pos < len(characters); pos++ {
		chr := characters[pos]
		data := font.Chars[int(chr)]
		for row, rowData := range data {
			for idx := 0; idx < charWidth; idx++ {
				mask := 1 << idx
				colour := dcled.Off
				if mask&int(rowData) == mask {
					colour = dcled.On
				}
				img.Set(pos*(charWidth+gap)+idx, row, colour)
			}
		}
	}
	return img
}

const (
	LoopHorizontally int = iota + 1
	LoopVertically
)

// Scroll will scroll an image across the device.
func Scroll(dev dcled.Device, img dcled.Canvas) {

	maxWidth := img.Bounds().Dx()
	x := 0
	dir := 1

	for {
		subimg := img.SubImage(image.Rect(x, 0, x+22, 8))
		_ = dcled.DisplayCanvas(subimg, dev)
		time.Sleep(50 * time.Millisecond)

		x += dir

		if x >= maxWidth {
			x = 0
		}
	}
}
