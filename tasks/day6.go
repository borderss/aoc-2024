package tasks

import (
	"fmt"
	"os"
	"strings"
)

type Day6 struct {
	data string
}

func (d *Day6) Init() error {
	val, err := os.ReadFile("./data/day6.txt")
	if err != nil {
		return err
	}

	d.data = string(val)
	return nil
}

func (d *Day6) Puzzle1() (any, error) {
	rows := strings.Split(d.data, "\n")

	mapData := make([][]rune, len(rows))
	for i, row := range rows {
		mapData[i] = []rune(row)
	}

	mapHeight := len(mapData)
	mapWidth := len(mapData[0])

	guardCoordinates := make([]int, 2)
	for y := 0; y < mapHeight; y++ {
		for x := 0; x < mapWidth; x++ {
			if mapData[y][x] == 94 {
				guardCoordinates = []int{x, y}
			}
		}
	}

	obstaclesCoordinates := make([][]int, 0)
	for y := 0; y < mapHeight; y++ {
		for x := 0; x < mapWidth; x++ {
			if mapData[y][x] == 35 {
				obstaclesCoordinates = append(obstaclesCoordinates, []int{x, y})
			}
		}
	}

	guardDirection := []int{0, -1}
	visitedLocations := make([][]int, 0)
	visitedLocations = append(visitedLocations, guardCoordinates)

	validMove := true
	for validMove {
		newGuardCoordinates, outOfBounds := moveGuard(mapWidth, mapHeight, guardCoordinates, guardDirection)
		fmt.Println(newGuardCoordinates, outOfBounds)

		if outOfBounds {
			validMove = false
			continue
		}

		if mapData[newGuardCoordinates[1]][newGuardCoordinates[0]] == 35 {
			if guardDirection[0] == 0 && guardDirection[1] == -1 {
				guardDirection = []int{1, 0}
			} else if guardDirection[0] == 1 && guardDirection[1] == 0 {
				guardDirection = []int{0, 1}
			} else if guardDirection[0] == 0 && guardDirection[1] == 1 {
				guardDirection = []int{-1, 0}
			} else if guardDirection[0] == -1 && guardDirection[1] == 0 {
				guardDirection = []int{0, -1}
			}

			continue
		}

		visitedLocations = appendIfUnique(visitedLocations, newGuardCoordinates)
		guardCoordinates = newGuardCoordinates
	}

	return len(visitedLocations), nil
}

type Location struct {
	Coordinate []int
	Direction  []int
}

func (d *Day6) Puzzle2() (any, error) {
	rows := strings.Split(d.data, "\n")

	mapData := make([][]rune, len(rows))
	for i, row := range rows {
		mapData[i] = []rune(row)
	}

	mapHeight := len(mapData)
	mapWidth := len(mapData[0])

	guardCoordinates := make([]int, 2)
	for y := 0; y < mapHeight; y++ {
		for x := 0; x < mapWidth; x++ {
			if mapData[y][x] == 94 {
				guardCoordinates = []int{x, y}
			}
		}
	}

	obstaclesCoordinates := make([][]int, 0)
	for y := 0; y < mapHeight; y++ {
		for x := 0; x < mapWidth; x++ {
			if mapData[y][x] == 35 {
				obstaclesCoordinates = append(obstaclesCoordinates, []int{x, y})
			}
		}
	}

	guardDirection := []int{0, -1}

	visitedLocations := make([]Location, 0)
	visitedLocations = append(visitedLocations, Location{guardCoordinates, guardDirection})

	isInInfiniteLoop := false

	validMove := true
	for validMove {
		newGuardCoordinates, outOfBounds := moveGuard(mapWidth, mapHeight, guardCoordinates, guardDirection)
		fmt.Println(newGuardCoordinates, outOfBounds)

		if outOfBounds {
			validMove = false
			continue
		}

		if mapData[newGuardCoordinates[1]][newGuardCoordinates[0]] == 35 {
			if guardDirection[0] == 0 && guardDirection[1] == -1 {
				guardDirection = []int{1, 0}
			} else if guardDirection[0] == 1 && guardDirection[1] == 0 {
				guardDirection = []int{0, 1}
			} else if guardDirection[0] == 0 && guardDirection[1] == 1 {
				guardDirection = []int{-1, 0}
			} else if guardDirection[0] == -1 && guardDirection[1] == 0 {
				guardDirection = []int{0, -1}
			}

			continue
		}

		if !isInInfiniteLoop {
			for i, loc := range visitedLocations {
				if loc.Coordinate[0] == newGuardCoordinates[0] && loc.Coordinate[1] == newGuardCoordinates[1] {
					isInInfiniteLoop = true
					visitedLocations = visitedLocations[:i]
					break
				}
			}
		}

		visitedLocations = appendIfUnique(visitedLocations, Location{newGuardCoordinates, guardDirection})
		guardCoordinates = newGuardCoordinates
	}

	return len(visitedLocations), nil
}

func moveGuard(mapx int, mapy int, guardCoordinates []int, guardDirection []int) ([]int, bool) {
	newX := guardCoordinates[0] + guardDirection[0]
	newY := guardCoordinates[1] + guardDirection[1]

	if isOutOfBounds(mapx, mapy, []int{newX, newY}) {
		return guardCoordinates, true
	}

	return []int{newX, newY}, false
}

func isOutOfBounds(mapX int, mapY int, coords []int) bool {
	if coords[0] < 0 || coords[0] >= mapX {
		return true
	}

	if coords[1] < 0 || coords[1] >= mapY {
		return true
	}

	return false
}

func appendIfUnique(slice []Location, element Location) []Location {
	for _, ele := range slice {
		if ele.Coordinate[0] == element.Coordinate[0] && ele.Coordinate[1] == element.Coordinate[1] {
			return slice
		}
	}

	return append(slice, element)
}
