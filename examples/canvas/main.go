package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/karalabe/hid"
	"github.com/redtoad/go-dcmb"
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

	var canvas dcmb.Canvas
	var ok bool
	if canvas, ok = img.(*image.NRGBA); !ok {
		panic("no NRGBA")
	}

	var list = hid.Enumerate(dcmb.VendorID, dcmb.ProductID)
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

	maxHeight := canvas.Bounds().Dy()
	y := 0
	dir := +1

	for {

		subimg := canvas.SubImage(image.Rect(0, y, 22, y+7))
		_ = dcmb.DisplayCanvas(subimg, device)
		time.Sleep(100 * time.Millisecond)

		if y+7 >= maxHeight {
			dir = -1
		}

		if y < 0 {
			dir = +1
		}

		y += dir
	}

}
