package main

import (
	"os"
	"log"
	"time"
	"fmt"
	"flag"
	"image"
	"image/png"
	_ "image/png"
	_ "image/jpeg"
	"github.com/esimov/stackblur-go"
	"github.com/fogleman/imview"
)

var (
	source		= flag.String("in", "", "Source")
	destination	= flag.String("out", "", "Destination")
	radius 		= flag.Int("radius", 20, "Radius")
)

func main() {
	flag.Parse()

	img, err := os.Open(*source)
	defer img.Close()

	src, _, err := image.Decode(img)
	if err != nil {
		panic(err)
	}
	start := time.Now()
	dst := stackblur.Process(src, uint32(src.Bounds().Dx()), uint32(src.Bounds().Dy()), uint32(*radius))
	end := time.Since(start)
	fmt.Printf("Processed in: %.2fs\n", end.Seconds())

	fq, err := os.Create(*destination)
	defer fq.Close()

	if err = png.Encode(fq, dst); err != nil {
		log.Fatal(err)
	}

	image, _ := imview.LoadImage(*destination)
	view := imview.ImageToRGBA(image)
	imview.Show(view)
}