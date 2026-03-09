package helpers

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func EncodeImage(rawData []byte) error {
	// padding
	for len(rawData)%3 != 0 {
		rawData = append(rawData, 0)
	}

	pixels := len(rawData) / 3
	side := int(math.Ceil(math.Sqrt(float64(pixels))))
	img := image.NewRGBA(image.Rect(0, 0, side, side))
	for i := 0; i < pixels; i++ {
		x := i % side
		y := i / side
		img.Set(x, y, color.RGBA{
			R: rawData[i*3],
			G: rawData[i*3+1],
			B: rawData[i*3+2],
			A: 255,
		})
	}
	f, err := os.Create("tmp.png")
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, img)
}
