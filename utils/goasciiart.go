// by mo2zie

package utils

import (
	"github.com/nfnt/resize"

	"bytes"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"reflect"
)

//   ,:;Il!i><_?}{)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhaoMW&%B@
var ASCIISTR = "  ,:I=!?I/XZONM&@"

func Init(width *int, path *string) (image.Image, int) {
	f, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	f.Close()
	return img, *width
}

func ScaleImage(img image.Image, w int) (image.Image, int, int) {
	sz := img.Bounds()
	h := (sz.Max.Y * w * 10) / (sz.Max.X * 16)
	img = resize.Resize(uint(w), uint(h), img, resize.Lanczos3)
	return img, w, h
}

func Convert2Ascii(img image.Image, w, h int) []byte {
	table := []byte(ASCIISTR)
	buf := new(bytes.Buffer)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			g := color.GrayModel.Convert(img.At(j, i))
			y := reflect.ValueOf(g).FieldByName("Y").Uint()
			pos := int(y * 16 / 255)
			_ = buf.WriteByte(table[pos])
		}
		_ = buf.WriteByte('\n')
	}
	return buf.Bytes()
}

func GetAscii(width int, img *image.Image) string {
	return string(Convert2Ascii(ScaleImage(*img, width)))
}
