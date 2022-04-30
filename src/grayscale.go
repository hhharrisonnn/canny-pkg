package canny

import (
	"image"
)

func Greyscale(img image.Image) [][]float64 {
	imgBound := img.Bounds()
	imgWidth, imgHeight := imgBound.Max.X, imgBound.Max.Y

	imageIndex := make([][]float64, imgWidth)
	for i := range imageIndex {
		imageIndex[i] = make([]float64, imgHeight)
	}

	for y := 0; y < imgHeight; y++ {
		for x := 0; x < imgWidth; x++ {
			imgColour := img.At(x, y)
			pixelGreyValue, _, _, _ := imgColour.RGBA()
			Y := uint8(pixelGreyValue)
			graycolour := float64(Y)
			imageIndex[x][y] = graycolour
		}
	}

	return imageIndex
}
