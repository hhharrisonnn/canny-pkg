package canny

import (
	"math"
)

func GaussianKernel(i, j int8, sigma float64) float64 {
	normalFunc := 1 / (2 * math.Pi * math.Pow(sigma, 2))

	xMat := [5][5]float64{
		{-2, -2, -2, -2, -2},
		{-1, -1, -1, -1, -1},
		{0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1},
		{2, 2, 2, 2, 2},
	}

	yMat := [5][5]float64{
		{-2, -1, 0, 1, 2},
		{-2, -1, 0, 1, 2},
		{-2, -1, 0, 1, 2},
		{-2, -1, 0, 1, 2},
		{-2, -1, 0, 1, 2},
	}

	xyMat := [5][5]float64{}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			xMatSquare := math.Pow(xMat[i][j], 2)

			yMatSquare := math.Pow(yMat[i][j], 2)

			xyMatDiv := -((xMatSquare + yMatSquare) / (2 * math.Pow(sigma, 2)))

			xyMatExp := math.Exp(xyMatDiv)

			xyMatFinal := xyMatExp * normalFunc

			xyMat[i][j] = xyMatFinal
		}
	}

	return xyMat[i+2][j+2]
}

func Gaussian(img [][]float64, sigma float64) [][]float64 {
	imageIndex := make([][]float64, len(img))
	for i := range imageIndex {
		imageIndex[i] = make([]float64, len(img[0]))
	}

	for j := 2; j < len(img[0])-2; j++ {
		for i := 2; i < len(img)-2; i++ {
			result := img[i-2][j+2]*GaussianKernel(-2, 2, sigma) +
				img[i-1][j+2]*GaussianKernel(-1, 2, sigma) +
				img[i][j+2]*GaussianKernel(0, 2, sigma) +
				img[i+1][j+2]*GaussianKernel(1, 2, sigma) +
				img[i+2][j+2]*GaussianKernel(2, 2, sigma) +
				img[i-2][j+1]*GaussianKernel(-2, 1, sigma) +
				img[i-1][j+1]*GaussianKernel(-1, 1, sigma) +
				img[i][j+1]*GaussianKernel(0, 1, sigma) +
				img[i+1][j+1]*GaussianKernel(1, 1, sigma) +
				img[i+2][j+1]*GaussianKernel(2, 1, sigma) +
				img[i-2][j]*GaussianKernel(-2, 0, sigma) +
				img[i-1][j]*GaussianKernel(-1, 0, sigma) +
				img[i][j]*GaussianKernel(0, 0, sigma) +
				img[i+1][j]*GaussianKernel(1, 0, sigma) +
				img[i+2][j]*GaussianKernel(2, 0, sigma) +
				img[i-2][j-1]*GaussianKernel(-2, -1, sigma) +
				img[i-1][j-1]*GaussianKernel(-1, -1, sigma) +
				img[i][j-1]*GaussianKernel(0, -1, sigma) +
				img[i+1][j-1]*GaussianKernel(1, -1, sigma) +
				img[i+2][j-1]*GaussianKernel(2, -1, sigma) +
				img[i-2][j-2]*GaussianKernel(-2, -2, sigma) +
				img[i-1][j-2]*GaussianKernel(-1, -2, sigma) +
				img[i][j-2]*GaussianKernel(0, -2, sigma) +
				img[i+1][j-2]*GaussianKernel(1, -2, sigma) +
				img[i+2][j-2]*GaussianKernel(2, -2, sigma)
			imageIndex[i][j] = result
		}
	}

	return imageIndex
}
