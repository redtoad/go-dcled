package main

import (
	"fmt"
	"time"

	"github.com/karalabe/hid"
	"github.com/redtoad/go-dcled"
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

	grid := [][]int{
		{1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0},
		{1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0},
		{1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1},
		{0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0},
	}

	for {
		_ = dcled.DisplayGrid(grid, device)
		time.Sleep(400 * time.Millisecond)
	}

}
