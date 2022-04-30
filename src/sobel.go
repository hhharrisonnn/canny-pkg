package canny

import (
	"math"
)

func sobelX(i, j int) float64 {
	xMat := [3][3]float64{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}

	return xMat[i][j]
}

func sobelY(i, j int) float64 {
	yMat := [3][3]float64{
		{-1, -2, -1},
		{0, 0, 0},
		{1, 2, 1},
	}

	return yMat[i][j]
}

func SobelConvolution(img [][]float64) ([][]float64, [][]float64) {
	imageIndex := make([][]float64, len(img))
	for i := range imageIndex {
		imageIndex[i] = make([]float64, len(img[0]))
	}

	theta := make([][]float64, len(img)-1)
	for i := range theta {
		theta[i] = make([]float64, len(img[0])-1)
	}

	for j := 1; j < len(img[0])-1; j++ {
		for i := 1; i < len(img)-1; i++ {
			Gx := img[i-1][j+1]*sobelX(0, 0) +
				img[i][j+1]*sobelX(1, 0) +
				img[i+1][j+1]*sobelX(2, 0) +
				img[i-1][j]*sobelX(0, 1) +
				img[i][j]*sobelX(1, 1) +
				img[i+1][j]*sobelX(2, 1) +
				img[i-1][j-1]*sobelX(0, 2) +
				img[i+1][j-1]*sobelX(1, 2) +
				img[i+1][j-1]*sobelX(2, 2)

			Gy := img[i-1][j+1]*sobelY(0, 0) +
				img[i][j+1]*sobelY(1, 0) +
				img[i+1][j+1]*sobelY(2, 0) +
				img[i-1][j]*sobelY(0, 1) +
				img[i][j]*sobelY(1, 1) +
				img[i+1][j]*sobelY(2, 1) +
				img[i-1][j-1]*sobelY(0, 2) +
				img[i+1][j-1]*sobelY(1, 2) +
				img[i+1][j-1]*sobelY(2, 2)

			imageIndex[i][j] = math.Abs(Gx) + math.Abs(Gy)

			theta[i][j] = math.Atan2(Gy, Gx) * 180 / math.Pi
		}
	}

	return imageIndex, theta
}

func Sobel(img [][]float64) ([][]float64, [][]float64) {
	imageIndex, theta := SobelConvolution(img)
	return imageIndex, theta
}
