package tasks

import (
	"fmt"
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
			a, errA := utils.ParseInt8(row[j])
			b, errB := utils.ParseInt8(row[j-1])
			if errA != nil {
				return nil, errA
			}
			if errB != nil {
				return nil, errB
			}

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
		rowLen := len(row)
		deltaArr := make([]int8, rowLen)

		for j := 1; j < rowLen; j++ {
			a, _ := utils.ParseInt8(row[j])
			b, _ := utils.ParseInt8(row[j-1])

			delta := a - b

			deltaArr[j] = delta
		}

		fmt.Printf("%v\n", deltaArr)

	}

	return safeRows, nil
}
