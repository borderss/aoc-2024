package tasks

import (
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

	guardLocation := Location{}
	for y := 0; y < mapHeight; y++ {
		for x := 0; x < mapWidth; x++ {
			if mapData[y][x] == 94 {
				guardLocation = Location{[]int{x, y}, []int{0, -1}}
			}
		}
	}

	visitedCoordinates := make([]Location, 0)
	visitedCoordinates = append(visitedCoordinates, guardLocation)

	validMove := true
	for validMove {
		newGuardLocation, outOfBounds := moveGuard(mapWidth, mapHeight, guardLocation)

		if outOfBounds {
			validMove = false
			continue
		}

		if mapData[newGuardLocation.Coordinate[1]][newGuardLocation.Coordinate[0]] == 35 {
			if guardLocation.Direction[0] == 0 && guardLocation.Direction[1] == -1 {
				guardLocation.Direction = []int{1, 0}
			} else if guardLocation.Direction[0] == 1 && guardLocation.Direction[1] == 0 {
				guardLocation.Direction = []int{0, 1}
			} else if guardLocation.Direction[0] == 0 && guardLocation.Direction[1] == 1 {
				guardLocation.Direction = []int{-1, 0}
			} else if guardLocation.Direction[0] == -1 && guardLocation.Direction[1] == 0 {
				guardLocation.Direction = []int{0, -1}
			}

			continue
		}

		visitedCoordinates = appendIfUniqueCoordinate(visitedCoordinates, newGuardLocation)
		guardLocation = newGuardLocation
	}

	return len(visitedCoordinates), nil
}

type Location struct {
	Coordinate []int
	Direction  []int
}

func (d *Day6) Puzzle2() (any, error) {
	// pilnīga miskaste; netaisos šo labot.

	rows := strings.Split(d.data, "\n")

	mapData := make([][]rune, len(rows))
	guard := Guard{
		directions: [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}},
	}

	for y, row := range rows {
		mapData[y] = make([]rune, len(row))
		for x, cell := range row {
			switch cell {
			case 94:
				guard.x = x
				guard.y = y
				mapData[y][x] = 46
			default:
				mapData[y][x] = cell
			}
		}
	}

	originalGuard := guard

	mapDataCopy := make([][]rune, len(mapData))
	for i := range mapDataCopy {
		mapDataCopy[i] = make([]rune, len(mapData[i]))
		copy(mapDataCopy[i], mapData[i])
	}

	for inBounds(guard.x, guard.y, mapData) {
		newX, newY := guard.nextPosition()

		if !inBounds(newX, newY, mapData) {
			guard.moveForward()
		} else if mapData[newY][newX] == 35 {
			guard.turnRight()
		} else {
			mapData[newY][newX] = 35
			if mapDataCopy[newY][newX] != 88 && causesLoop(originalGuard, mapData) {
				mapDataCopy[newY][newX] = 88
			}
			mapData[newY][newX] = 46
			guard.moveForward()
		}
	}

	sum := 0
	for y := range mapDataCopy {
		for x := range mapDataCopy[y] {
			if mapDataCopy[y][x] == 88 {
				sum++
			}
		}
	}

	return sum, nil
}

type Guard struct {
	x, y       int
	dirIndex   int
	directions [4][2]int
}

func (g *Guard) nextPosition() (int, int) {
	return g.x + g.directions[g.dirIndex][0], g.y + g.directions[g.dirIndex][1]
}

func (g *Guard) turnRight() {
	g.dirIndex = (g.dirIndex + 1) % 4
}

func (g *Guard) moveForward() {
	g.x, g.y = g.nextPosition()
}

func inBounds(x, y int, grid [][]rune) bool {
	return x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid)
}

func causesLoop(g Guard, grid [][]rune) bool {
	clone := g
	walls := map[[3]int]struct{}{}

	for inBounds(clone.x, clone.y, grid) {
		newX, newY := clone.nextPosition()

		if !inBounds(newX, newY, grid) {
			clone.moveForward()
		} else if grid[newY][newX] == 35 {
			wallKey := [3]int{newX, newY, clone.dirIndex}
			if _, exists := walls[wallKey]; exists {
				return true
			}
			walls[wallKey] = struct{}{}
			clone.turnRight()
		} else {
			clone.moveForward()
		}
	}

	return false
}

func moveGuard(mapx int, mapy int, location Location) (Location, bool) {
	newX := location.Coordinate[0] + location.Direction[0]
	newY := location.Coordinate[1] + location.Direction[1]

	if newX < 0 || newX >= mapx || newY < 0 || newY >= mapy {
		return location, true
	}

	return Location{
		Coordinate: []int{newX, newY},
		Direction:  location.Direction,
	}, false
}

func appendIfUniqueCoordinate(slice []Location, element Location) []Location {
	for _, ele := range slice {
		if ele.Coordinate[0] == element.Coordinate[0] && ele.Coordinate[1] == element.Coordinate[1] {
			return slice
		}
	}

	return append(slice, element)
}
