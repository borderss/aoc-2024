package tasks

import "errors"

type Day1 struct {
	data string
}

func (d *Day1) Init() error {
	d.data = "day1 data"
	return nil
}

func (d *Day1) Puzzle1() (string, error) {
	return d.data + " puzzle1", nil
}

func (d *Day1) Puzzle2() (string, error) {
	return "", errors.New("not implemented")
}
