package tasks

import (
	"os"
	"strings"

	"github.com/borderss/aoc-2024/utils"
)

type Day2 struct {
	data string
}

func (d *Day2) Init() error {
	val, err := os.ReadFile("./data/day2.txt")
	if err != nil {
		return err
	}

	d.data = string(val)
	return nil
}

func (d *Day2) Puzzle1() (any, error) {
	safeRows := 0
	rows := strings.Split(d.data, "\n")

	for i := 0; i < len(rows); i++ {
		var asc int
		row := strings.Fields(rows[i])
		rowLen := len(row)

		for j := 1; j < rowLen; j++ {
			a := utils.ParseInt8(row[j])
			b := utils.ParseInt8(row[j-1])

			delta := a - b

			if delta < 0 {
				if asc == 0 {
					asc = -1
				} else if asc == 1 {
					break
				}

				delta = -delta
			} else {
				if asc == 0 {
					asc = 1
				} else if asc == -1 {
					break
				}
			}

			if delta < 1 || delta > 3 {
				break
			}

			if j+1 == rowLen {
				safeRows++
			}
		}
	}

	return safeRows, nil
}

func (d *Day2) Puzzle2() (any, error) {
	safeRows := 0
	rows := strings.Split(d.data, "\n")

	for i := 0; i < len(rows); i++ {
		row := strings.Fields(rows[i])

		failIndex := findUnsafeIndex(row)
		if failIndex == -1 {
			safeRows++
			continue
		}

		// bruteforce, man manam inputam distance starp kļūdas punktu un maināmo elementu lai būtu ok varēja būt > 3, šā vai tā sanāktu bruteforce
		isSafe := false
		for j := 0; j < len(row); j++ {
			res := append(make([]string, 0, len(row)-1), row[:j]...)
			res = append(res, row[j+1:]...)

			if findUnsafeIndex(res) == -1 {
				isSafe = true
				break
			}
		}

		if isSafe {
			safeRows++
		}
	}

	return safeRows, nil
}

func findUnsafeIndex(row []string) int {
	rowLen := len(row)
	if rowLen < 2 {
		return -1
	}

	dir := utils.ParseInt8(row[1]) > utils.ParseInt8(row[0])

	for i := 1; i < rowLen; i++ {
		a := utils.ParseInt8(row[i])
		b := utils.ParseInt8(row[i-1])

		delta := a - b

		if dir && delta < 0 {
			return i
		} else if !dir && (delta > 0) {
			return i - 1
		} else if delta == 0 {
			return i
		}

		if delta < 0 {
			delta = -delta
		}

		if delta < 1 || delta > 3 {
			return i
		}
	}

	return -1
}
