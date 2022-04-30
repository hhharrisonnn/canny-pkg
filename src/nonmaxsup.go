package canny

func NonMaxSuppression(img [][]float64, theta [][]float64) [][]float64 {
	for j := 1; j < len(img[0])-1; j++ {
		for i := 1; i < len(img)-1; i++ {
			angle := theta[i][j]

			var a int
			var b int

			if (0 <= angle && angle < 180) || (157.5 <= angle && angle < 180) {
				a = int(img[i][j+1])
				b = int(img[i][j-1])
			} else if 22.5 <= angle && angle < 67.5 {
				a = int(img[i+1][j-1])
				b = int(img[i-1][j+1])
			} else if 67.5 <= angle && angle < 112.5 {
				a = int(img[i+1][j])
				b = int(img[i-1][j])
			} else if 112.5 <= angle && angle < 157.5 {
				a = int(img[i-1][j-1])
				b = int(img[i+1][j+1])
			}

			if (int(img[i][j]) >= a) && (int(img[i][j])) >= b {
				img[i][j] *= 1
			} else {
				img[i][j] = 0
			}
		}
	}

	return img
}