package canny

func DoubleThreshold(img [][]float64, highThresholdRatio float64, lowThresholdRatio float64) [][]float64 {
	var strong float64 = 255
	var weak float64 = 100

	highThreshold := strong * highThresholdRatio
	lowThreshold := highThreshold * lowThresholdRatio

	for j := 1; j < len(img[0])-1; j++ {
		for i := 1; i < len(img)-1; i++ {
			if img[i][j] >= highThreshold {
				img[i][j] = strong
			}
			if (img[i][j] < highThreshold) && (img[i][j] > lowThreshold) {
				img[i][j] = weak
			}
			if img[i][j] <= lowThreshold {
				img[i][j] = 0
			}
		}
	}

	return img
}
