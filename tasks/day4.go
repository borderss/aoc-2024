package tasks

import (
	"os"
	"strings"
)

type Day4 struct {
	data string
}

func (d *Day4) Init() error {
	val, err := os.ReadFile("./data/day4.txt")
	if err != nil {
		return err
	}

	d.data = string(val)
	return nil
}

func (d *Day4) Puzzle1() (any, error) {
	rows := strings.Split(d.data, "\n")
	lenX := len(rows[0])
	lenY := len(rows)

	grid := make([][]rune, lenY)
	for i := range grid {
		grid[i] = make([]rune, lenX)
	}

	for i, row := range rows {
		for j, r := range row {
			grid[i][j] = r
		}
	}

	foundWordCoords := 0

	for i := 0; i < lenY; i++ {
		for j := 0; j < lenX; j++ {
			if grid[i][j] == 'X' {
				foundWordCoords += findWordsFromCoordinate2D(grid, i, j)
			}
		}
	}

	return foundWordCoords, nil
}

func (d *Day4) Puzzle2() (any, error) {
	rows := strings.Split(d.data, "\n")
	lenX := len(rows[0])
	lenY := len(rows)

	grid := make([][]rune, lenY)
	for i := range grid {
		grid[i] = make([]rune, lenX)
	}

	for i, row := range rows {
		for j, r := range row {
			grid[i][j] = r
		}
	}

	foundWords := 0

	for i := 0; i < lenY; i++ {
		for j := 0; j < lenX; j++ {
			if grid[i][j] == 'A' {
				foundWords += findMasCrossWordsFromCoordinate2D(grid, i, j)
			}
		}
	}

	return foundWords, nil
}

func findWordsFromCoordinate2D(grid [][]rune, x, y int) int {
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {1, -1}, {-1, -1}, {-1, 1}}
	words := 0

	for _, dir := range directions {
		xDir, yDir := dir[0], dir[1]
		xNext, yNext := x+xDir, y+yDir
		if xNext >= 0 && xNext < len(grid) && yNext >= 0 && yNext < len(grid[0]) {
			if grid[xNext][yNext] == 'M' {
				xNext, yNext = xNext+xDir, yNext+yDir
				if xNext >= 0 && xNext < len(grid) && yNext >= 0 && yNext < len(grid[0]) {
					if grid[xNext][yNext] == 'A' {
						xNext, yNext = xNext+xDir, yNext+yDir
						if xNext >= 0 && xNext < len(grid) && yNext >= 0 && yNext < len(grid[0]) {
							if grid[xNext][yNext] == 'S' {
								words++
							}
						}
					}
				}
			}
		}
	}

	return words
}

func findMasCrossWordsFromCoordinate2D(grid [][]rune, x, y int) int {
	var foundDiagonal bool

	xTopLeft, yTopLeft := x-1, y-1
	xBottomRight, yBottomRight := x+1, y+1
	if xTopLeft >= 0 && yTopLeft >= 0 && xBottomRight < len(grid) && yBottomRight < len(grid[0]) {
		if grid[xTopLeft][yTopLeft] == 'M' && grid[xBottomRight][yBottomRight] == 'S' {
			foundDiagonal = true
		} else if grid[xTopLeft][yTopLeft] == 'S' && grid[xBottomRight][yBottomRight] == 'M' {
			foundDiagonal = true
		} else {
			return 0
		}
	}

	if !foundDiagonal {
		return 0
	}

	xTopRight, yTopRight := x-1, y+1
	xBottomLeft, yBottomLeft := x+1, y-1
	if xTopRight >= 0 && yTopRight < len(grid[0]) && xBottomLeft < len(grid) && yBottomLeft >= 0 {
		if grid[xTopRight][yTopRight] == 'M' && grid[xBottomLeft][yBottomLeft] == 'S' {
			return 1
		} else if grid[xTopRight][yTopRight] == 'S' && grid[xBottomLeft][yBottomLeft] == 'M' {
			return 1
		} else {
			return 0
		}
	}

	return 0
}
