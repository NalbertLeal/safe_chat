package main

import (
  "os"
	// "bytes"
	"fmt"
	// "image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"log"
  "bufio"
)

func drawablePNGImage(r io.Reader) (draw.Image, error) {
	img, err := png.Decode(bufio.NewReader(r))
	if err != nil {
		return nil, err
	}
	dimg, ok := img.(draw.Image)
	if !ok {
		return nil, fmt.Errorf("%T is not a drawable image type", img)
	}
	return dimg, nil
}

func main() {
	// test
	// buf := new(bytes.Buffer)
	// i := image.NewRGBA(image.Rect(0, 0, 10, 10))
	// if err := png.Encode(buf, i); err != nil {
	// 	log.Fatal(err)
	// }

  file, _ := os.Open("imgTest.png")

	img, err := drawablePNGImage(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Model:", img.ColorModel())
	fmt.Println("Bounds:", img.Bounds())
	fmt.Println("At(1,2):", img.At(1, 2))
	img.Set(1, 2, color.White)
	fmt.Println("At(1,2):", img.At(1, 2), "(after Set)")
}
