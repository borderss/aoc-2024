package tasks

import (
	"os"
	"strings"
)

type Day8 struct {
	data string
}

func (d *Day8) Init() error {
	val, err := os.ReadFile("./data/day8.txt")
	if err != nil {
		return err
	}

	d.data = string(val)
	return nil
}

func (d *Day8) Puzzle1() (any, error) {
	rows := strings.Split(d.data, "\n")
	mapData := make(map[rune][][]int)
	maxY := len(rows)
	maxX := len(rows[0])

	for y, row := range rows {
		for x, cell := range row {
			if cell == '.' {
				continue
			}

			mapData[cell] = append(mapData[cell], []int{x, y})
		}
	}

	allSignals := make([][]int, 0)

	for _, points := range mapData {
		if len(points) > 1 {
			for i := 0; i < len(points); i++ {
				for j := i + 1; j < len(points); j++ {
					signalPoints := makeSignalPoints(points[i], points[j], maxX, maxY)

					for _, signal := range signalPoints {
						allSignals = appendIfUnique(allSignals, signal)
					}
				}
			}
		}
	}

	return len(allSignals), nil
}

func (d *Day8) Puzzle2() (any, error) {
	rows := strings.Split(d.data, "\n")
	mapData := make(map[rune][][]int)
	maxY := len(rows)
	maxX := len(rows[0])

	for y, row := range rows {
		for x, cell := range row {
			if cell == '.' {
				continue
			}

			mapData[cell] = append(mapData[cell], []int{x, y})
		}
	}

	allSignals := make([][]int, 0)

	for _, points := range mapData {
		if len(points) > 1 {
			for i := 0; i < len(points); i++ {
				for j := i + 1; j < len(points); j++ {
					signalPoints := makeExtendedSignalPoints(points[i], points[j], maxX, maxY)

					for _, signal := range signalPoints {
						allSignals = appendIfUnique(allSignals, signal)
					}
				}
			}
		}
	}

	return len(allSignals), nil
}

func makeSignalPoints(src1 []int, src2 []int, maxX int, maxY int) [][]int {
	var result [][]int

	point1 := make([]int, 2)
	point2 := make([]int, 2)

	deltaX := src2[0] - src1[0]
	deltaY := src2[1] - src1[1]

	point1[0] = src1[0] - deltaX
	point1[1] = src1[1] - deltaY

	point2[0] = src2[0] + deltaX
	point2[1] = src2[1] + deltaY

	if maxX > point1[0] && point1[0] >= 0 && maxY > point1[1] && point1[1] >= 0 {
		result = append(result, point1)
	}

	if maxX > point2[0] && point2[0] >= 0 && maxY > point2[1] && point2[1] >= 0 {
		result = append(result, point2)
	}

	return result
}

func makeExtendedSignalPoints(src1 []int, src2 []int, maxX int, maxY int) [][]int {
	var result [][]int

	deltaX := src2[0] - src1[0]
	deltaY := src2[1] - src1[1]

	x, y := src1[0], src1[1]
	for {
		x -= deltaX
		y -= deltaY
		if x < 0 || x >= maxX || y < 0 || y >= maxY {
			break
		}
		result = append(result, []int{x, y})
	}

	x, y = src2[0], src2[1]
	for {
		x += deltaX
		y += deltaY
		if x < 0 || x >= maxX || y < 0 || y >= maxY {
			break
		}
		result = append(result, []int{x, y})
	}

	result = append(result, src1)
	result = append(result, src2)

	return result
}

func appendIfUnique(signals [][]int, signal []int) [][]int {
	for _, s := range signals {
		if s[0] == signal[0] && s[1] == signal[1] {
			return signals
		}
	}

	return append(signals, signal)
}
