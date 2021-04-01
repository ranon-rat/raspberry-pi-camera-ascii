package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"time"

	"github.com/dhowden/raspicam"
)

var letters = []string{}

func addthis() {
	pos := " .:!|l1G0@"
	for _, x := range pos {
		for y := 0; y < 257/(len(pos)); y++ {
			letters = append(letters, string(x))
		}
	}
}

func openThis(f io.Reader) {
	//this open the image and print the pixels
	img, _, _ := image.Decode(f)
	division := 27
	limitY, limitX := img.Bounds().Max.Y/division, img.Bounds().Max.X/division

	for x := img.Bounds().Min.X; x < limitX; x++ {
		for y := img.Bounds().Min.Y; y < limitY; y++ {
			r, g, b, _ := img.At(x*division, y*division).RGBA()
			fmt.Print(letters[int(((r/257)+(g/257)+(b/257))/3)%len(letters)])
		}
		fmt.Println()

	}
}

func main() {
	addthis()
	s := raspicam.NewStill()
	for {
		buffer := new(bytes.Buffer)
		errCh := make(chan error)
		/*go func() {
			for x := range errCh {
				fmt.Fprintf(os.Stderr, "%v\n", x)
			}
		}()*/
		raspicam.Capture(s, buffer, errCh)
		go openThis(buffer)
	}

}
