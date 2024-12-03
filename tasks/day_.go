package tasks

import (
	"os"
)

type Day_ struct {
	data string
}

func (d *Day_) Init() error {
	val, err := os.ReadFile("./data/day_.txt")
	if err != nil {
		return err
	}

	d.data = string(val)
	return nil
}

func (d *Day_) Puzzle1() (any, error) {

	return 0, nil
}

func (d *Day_) Puzzle2() (any, error) {

	return 0, nil
}
