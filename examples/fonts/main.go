package main

import (
	"fmt"
	"image"

	"github.com/karalabe/hid"
	"github.com/redtoad/go-dcled"
	"github.com/redtoad/go-dcled/fonts"
)

func main() {

	img := fonts.Text("Hallo Margit! Wollen wir ins Bett?", fonts.SmallInvFont)

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

	fonts.Scroll(device, img.(*image.NRGBA))

}
