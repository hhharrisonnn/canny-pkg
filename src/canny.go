package canny

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"os/user"
)

func Canny(fi string, sigma float64, highThresholdRatio float64, lowThresholdRatio float64) {
	inputImg, err := os.Open(fi)
	if err != nil {
		fmt.Printf("Failed to open %s: %s", fi, err)
		panic(err.Error())
	}
	defer inputImg.Close()

	img, _, err := image.Decode(inputImg)
	if err != nil {
		panic(err.Error())
	}

	greyscaleImg := Greyscale(img)
	gaussianImg := Gaussian(greyscaleImg, sigma)
	sobelImg, theta := Sobel(gaussianImg)
	nonMaxImg := NonMaxSuppression(sobelImg, theta)
	doubleThresholdImg := DoubleThreshold(nonMaxImg, highThresholdRatio, lowThresholdRatio)
	hysteresis := Hysteresis(doubleThresholdImg)

	userPath, _ := user.Current()
	newFi, err := os.Create(userPath.HomeDir + "/canny-pkg/img/newImage.png")
	if err != nil {
		fmt.Printf("Failed to create %s: %s", newFi, err)
		panic(err.Error())
	}
	defer newFi.Close()
	png.Encode(newFi, hysteresis)
}
