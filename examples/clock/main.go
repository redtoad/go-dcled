package main

import (
	"fmt"
	"time"

	"github.com/karalabe/hid"
	"github.com/redtoad/go-dcled"
	"github.com/redtoad/go-dcled/fonts"
)

func main() {

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

	for {
		now := time.Now()
		img := fonts.Text(fmt.Sprintf(">>>> %02d:%02d <<<<", now.Hour(), now.Minute()), fonts.SmallFont)
		dcled.Center(device, img)
		time.Sleep(dcled.MinimumRefreshRate)
	}

}
