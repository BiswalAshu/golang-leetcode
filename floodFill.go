package main

func fill(image [][]int, row int, col int, colour int, originalColour int) [][]int {
	if row < 0 || col < 0 {
		return image
	}
	if row >= len(image) || col >= len(image[row]) || image[row][col] == colour {
		return image
	}

	if image[row][col] == originalColour {
		image[row][col] = colour
		// right
		image = fill(image, row+1, col, colour, originalColour)
		// left
		image = fill(image, row-1, col, colour, originalColour)
		// top
		image = fill(image, row, col+1, colour, originalColour)
		// botom
		image = fill(image, row, col-1, colour, originalColour)
	}
	return image

}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	originalColour := image[sr][sc]
	fill(image, sr, sc, color, originalColour)
	return image
}
