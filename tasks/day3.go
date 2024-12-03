package tasks

import (
	"os"
	"regexp"

	"github.com/borderss/aoc-2024/utils"
)

type Day3 struct {
	data string
}

func (d *Day3) Init() error {
	val, err := os.ReadFile("./data/day3.txt")
	if err != nil {
		return err
	}

	d.data = string(val)
	return nil
}

func (d *Day3) Puzzle1() (any, error) {
	var sum int32

	r, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)

	for _, v := range r.FindAllStringSubmatch(d.data, -1) {
		sum += utils.ParseInt[int32](v[1]) * utils.ParseInt[int32](v[2])
	}

	return sum, nil
}

func (d *Day3) Puzzle2() (any, error) {
	var sum int32

	r, _ := regexp.Compile(`do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\)`)

	include := true
	for _, v := range r.FindAllStringSubmatch(d.data, -1) {
		switch v[0] {
		case "don't()":
			include = false
		case "do()":
			include = true
		default:
			if include {
				sum += utils.ParseInt[int32](v[1]) * utils.ParseInt[int32](v[2])
			}
		}
	}

	return sum, nil
}
