
# go-dcled

Go library for interacting with Dream Cheeky Message Board

## Usage

```go
package main

import (
  "fmt"
  "time"

  "github.com/karalabe/hid"
  "github.com/redtoad/go-dcled"
)

func main() {
	var list = hid.Enumerate(0x1D34, 0x0013)
	// Use first device
	var device, err = list[0].Open()
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
``` 


Based on code by [MylesIsCool](https://gist.github.com/MylesIsCool/227a64a679fb0fc8432fe1c342f526dd). 
Very helpful: [The Last Outpost ](https://www.last-outpost.com/~malakai/dcled/)