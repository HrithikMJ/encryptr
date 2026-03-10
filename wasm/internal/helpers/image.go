package helpers

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
)

func EncodeImage(rawData []byte) (rawImage []byte, err error) {
	// Prepend 4-byte big-endian length so we can recover exact size on decode
	dataLen := make([]byte, 4)
	binary.BigEndian.PutUint32(dataLen, uint32(len(rawData)))
	rawData = append(dataLen, rawData...)

	// padding to make length divisible by 3 (one pixel = 3 bytes)
	for len(rawData)%3 != 0 {
		rawData = append(rawData, 0)
	}
	pixels := len(rawData) / 3
	side := int(math.Ceil(math.Sqrt(float64(pixels))))

	img := image.NewRGBA(image.Rect(0, 0, side, side))
	for i := range rawData {
		x := i % side
		y := i / side
		img.Set(x, y, color.RGBA{
			R: rawData[i*3],
			G: rawData[i*3+1],
			B: rawData[i*3+2],
			A: 255,
		})
	}
	buf := bytes.NewBuffer(nil)
	if err = png.Encode(buf, img); err != nil {
		return nil, err
	}
	rawImage = buf.Bytes()
	return
}

func DecodeImage(rawImage []byte) ([]byte, error) {
	img, err := png.Decode(bytes.NewReader(rawImage))
	if err != nil {
		return nil, err
	}
	bounds := img.Bounds()
	data := make([]byte, 0, bounds.Dx()*bounds.Dy()*3)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()

			data = append(data, byte(r>>8), byte(g>>8), byte(b>>8))
		}
	}
	// fmt.Println("img:", len(data))

	// use this to recover the original data length
	if len(data) < 4 {
		return nil, fmt.Errorf("image data too short to contain length header")
	}
	originalLen := binary.BigEndian.Uint32(data[:4])
	if int(originalLen) > len(data)-4 {
		return nil, fmt.Errorf("invalid data length in header: %d (available: %d)", originalLen, len(data)-4)
	}
	// fmt.Println("originalLen:", originalLen)
	return data[4 : 4+originalLen], nil
}
