package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"time"

	"github.com/karalabe/hid"
	"github.com/redtoad/go-dcled"
)

func main() {

	f, err := os.Open("gopher2.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		// replace this with real error handling
		panic(err)
	}

	var canvas draw.Image
	var ok bool
	if canvas, ok = img.(*image.NRGBA); !ok {
		panic("no NRGBA")
	}

	var list = hid.Enumerate(dcled.VendorID, dcled.ProductID)
	if len(list) == 0 {
		println("Could not find USB device! Is it plugged in?")
		return
	}

	// Use first device
	device, err := list[0].Open()
	if err != nil {
		panic(err)
	}

	println(fmt.Sprintf("Connected to %s %s", device.Manufacturer, device.Product))

	y := 0
	dir := +1

	dst := dcled.NewCanvas()
	sr := canvas.Bounds()

	for {
		r := sr.Sub(image.Point{0, y})
		draw.Draw(dst, r, img, sr.Min, draw.Src)
		_ = dcled.DisplayCanvas(dst, device)
		time.Sleep(50 * time.Millisecond)

		if y+7 >= sr.Dy() {
			dir = -1
		}

		if y < 0 {
			dir = +1
		}

		y += dir
	}

}
