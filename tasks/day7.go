package tasks

import (
	"fmt"
	"os"
	"strings"

	"github.com/borderss/aoc-2024/utils"
)

type Day7 struct {
	data string
}

func (d *Day7) Init() error {
	val, err := os.ReadFile("./data/day7.txt")
	if err != nil {
		return err
	}

	d.data = string(val)
	return nil
}

func (d *Day7) Puzzle1() (any, error) {
	rows := strings.Split(d.data, "\n")

	var sum int64
	for _, row := range rows {
		split := strings.Split(row, ": ")
		targetValue := utils.ParseInt[int64](split[0])
		numbersRaw := strings.Split(split[1], " ")
		numbers := make([]int64, 0, len(numbersRaw))
		for _, number := range numbersRaw {
			numbers = append(numbers, utils.ParseInt[int64](number))
		}

		checked := make(map[string]bool)

		var depthFirstSearch func(index int, sum int64) bool
		depthFirstSearch = func(index int, sum int64) bool {
			if index == len(numbers) {
				return sum == targetValue
			}

			if sum > targetValue {
				return false
			}

			add := depthFirstSearch(index+1, sum+numbers[index])
			mult := depthFirstSearch(index+1, sum*numbers[index])

			checked[fmt.Sprintf("%d:%d", index, sum)] = add || mult
			return add || mult
		}

		if depthFirstSearch(1, numbers[0]) {
			sum += targetValue
		}
	}

	return sum, nil
}

func (d *Day7) Puzzle2() (any, error) {
	rows := strings.Split(d.data, "\n")

	var sum int64
	for _, row := range rows {
		split := strings.Split(row, ": ")
		targetValue := utils.ParseInt[int64](split[0])
		numbersRaw := strings.Split(split[1], " ")
		numbers := make([]int64, 0, len(numbersRaw))
		for _, number := range numbersRaw {
			numbers = append(numbers, utils.ParseInt[int64](number))
		}

		checked := make(map[string]bool)

		var depthFirstSearch func(index int, sum int64) bool
		depthFirstSearch = func(index int, sum int64) bool {
			if index == len(numbers) {
				return sum == targetValue
			}

			if sum > targetValue {
				return false
			}

			add := depthFirstSearch(index+1, sum+numbers[index])
			mult := depthFirstSearch(index+1, sum*numbers[index])
			concat := depthFirstSearch(index+1, utils.ParseInt[int64](fmt.Sprintf("%d%d", sum, numbers[index])))

			checked[fmt.Sprintf("%d:%d", index, sum)] = add || mult || concat
			return add || mult || concat
		}

		if depthFirstSearch(1, numbers[0]) {
			sum += targetValue
		}
	}

	return sum, nil
}
