package main

import (
	"encoding/binary"
	"fmt"
	"github.com/karalabe/hid"
	"time"
)

func main() {
	var list = hid.Enumerate(0x1D34, 0x0013)
	if len(list) == 0 {
		println("Could not find USB device! Is it plugged in?")
		return
	}

	// Use first device
	var device, err = list[0].Open()
	if err != nil {
		panic(err)
	}

	println(fmt.Sprintf("Connected to %s %s", device.Manufacturer, device.Product))

	//grid := [][]int{
	//	{1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0},
	//	{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1},
	//	{1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0},
	//	{1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0},
	//	{1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0},
	//	{1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1},
	//	{0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0},
	//}

	for {
		//	_ = DisplayGrid(grid, device)
		//	time.Sleep(400 * time.Millisecond)

		//////
		for _, row := range []uint64{0x2001ffefffffd7f, 0x2021ffbbffff7df, 0x2041ffbbffffd7f, 0x2061ffeff000000} {
			buffer := make([]byte, 8)
			binary.BigEndian.PutUint64(buffer, row)
			_, _ = device.Write(buffer)
		}
		println(1)
		time.Sleep(400 * time.Millisecond)

		for _, row := range []uint64{0x2001ffeff1ffd7f, 0x2021ffbbf1ff7df, 0x2041ffbbf1ffd7f, 0x2061ffeff000000} {
			buffer := make([]byte, 8)
			binary.BigEndian.PutUint64(buffer, row)
			_, _ = device.Write(buffer)
		}
		println(2)
		time.Sleep(400 * time.Millisecond)

		_ = DisplayGrid([][]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		}, device)
		println(3)
		time.Sleep(400 * time.Millisecond)

	}

}
