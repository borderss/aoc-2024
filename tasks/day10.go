package tasks

import (
	"os"
	"strings"
)

type Day10 struct {
	data string
}

func (d *Day10) Init() error {
	val, err := os.ReadFile("./data/day10.txt")
	if err != nil {
		return err
	}

	d.data = string(val)
	return nil
}

func (d *Day10) Puzzle1() (any, error) {
	grid := strings.Split(d.data, "\n")
	height := len(grid)
	width := len(grid[0])
	sum := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == '0' {
				sum += countUniquePaths(determinePath(grid, y, x))
			}
		}
	}

	return sum, nil
}

func (d *Day10) Puzzle2() (any, error) {
	grid := strings.Split(d.data, "\n")
	height := len(grid)
	width := len(grid[0])
	sum := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == '0' {
				sum += determinePathCount(grid, y, x)
			}
		}
	}

	return sum, nil
}

func findNextGrids(grid []string, y, x int) [][2]int {
	targetValue := grid[y][x] - 47
	foundGrids := make([][2]int, 0)

	if y > 0 && grid[y-1][x]-48 == targetValue {
		foundGrids = append(foundGrids, [2]int{y - 1, x})
	}
	if y+1 < len(grid) && grid[y+1][x]-48 == targetValue {
		foundGrids = append(foundGrids, [2]int{y + 1, x})
	}
	if x > 0 && grid[y][x-1]-48 == targetValue {
		foundGrids = append(foundGrids, [2]int{y, x - 1})
	}
	if x+1 < len(grid[0]) && grid[y][x+1]-48 == targetValue {
		foundGrids = append(foundGrids, [2]int{y, x + 1})
	}

	return foundGrids
}

func determinePath(grid []string, y, x int) [][2]int {
	if grid[y][x] == '9' {
		return [][2]int{{y, x}}
	}

	if grid[y][x] == '0' {
		nextSteps := findNextGrids(grid, y, x)
		totalPaths := [][2]int{}
		for _, step := range nextSteps {
			totalPaths = append(totalPaths, determinePath(grid, step[0], step[1])...)
		}

		return totalPaths
	}

	nextSteps := findNextGrids(grid, y, x)
	totalPaths := [][2]int{}
	for _, step := range nextSteps {
		totalPaths = append(totalPaths, determinePath(grid, step[0], step[1])...)
	}

	return totalPaths
}

func determinePathCount(grid []string, y, x int) int {
	if grid[y][x] == '9' {
		return 1
	}

	if grid[y][x] == '0' {
		nextSteps := findNextGrids(grid, y, x)
		totalPaths := 0
		for _, step := range nextSteps {
			totalPaths += determinePathCount(grid, step[0], step[1])
		}

		return totalPaths
	}

	nextSteps := findNextGrids(grid, y, x)
	totalPaths := 0
	for _, step := range nextSteps {
		totalPaths += determinePathCount(grid, step[0], step[1])
	}

	return totalPaths
}

func countUniquePaths(paths [][2]int) int {
	uniquePaths := make(map[[2]int]bool)
	for _, path := range paths {
		uniquePaths[path] = true
	}

	return len(uniquePaths)
}
