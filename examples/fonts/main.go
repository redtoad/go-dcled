package main

import (
	"fmt"
	"image"

	"github.com/karalabe/hid"
	"github.com/redtoad/go-dcled"
	"github.com/redtoad/go-dcled/fonts"
)

func main() {

	img := fonts.Text("The quick fox jumped over the lazy fox! @#§\"%&?/()", fonts.SmallInvFont)

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

	dcled.Scroll(device, img.(*image.NRGBA))

}
