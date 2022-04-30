package canny

import (
	"image"
	"image/color"
)

func Hysteresis(img [][]float64) *image.Gray {
	var strong float64 = 255
	var weak float64 = 100

	newImage := image.NewGray((image.Rectangle{image.Point{1, 1}, image.Point{len(img) - 1, len(img[0]) - 1}}))

	for j := 1; j < len(img[0])-1; j++ {
		for i := 1; i < len(img)-1; i++ {
			if img[i][j] == weak {
				if (img[i-1][j+1] == strong) || (img[i][j+1] == strong) || (img[i+1][j+1] == strong) ||
					(img[i-1][j] == strong) || (img[i][j] == strong) || (img[i+1][j] == strong) ||
					(img[i-1][j-1] == strong) || (img[i][j-1] == strong) || (img[i+1][j-1] == strong) {
					img[i][j] = strong
				} else {
					img[i][j] = 0
				}
			}
			graycolour := color.Gray{uint8(img[i][j])}
			newImage.Set(i, j, graycolour)
		}
	}

	return newImage
}
