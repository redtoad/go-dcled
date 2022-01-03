package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/redtoad/go-dcled/fonts"
)

func main() {

	for _, font := range []fonts.Font{
		fonts.DefaultFont,
		fonts.SgaFont,
		fonts.SmallInvFont,
		fonts.XxxFont,
	} {

		txt := ""
		for c := range font.Chars {
			txt += string(rune(c))
		}

		img := fonts.Text(txt, font)

		f, err := os.Create(fmt.Sprintf("font_%s.png", font.Name))
		if err != nil {
			log.Fatalf("could not create file: %v", err)
		}
		defer f.Close()

		// Encode to `PNG` with `DefaultCompression` level
		// then save to file
		err = png.Encode(f, img)
		if err != nil {
			log.Fatalf("could not write image: %v", err)
		}
	}
}
