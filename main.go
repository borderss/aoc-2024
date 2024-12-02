package main

import (
	"github.com/borderss/aoc-2024/tasks"
)

func main() {
	task := &tasks.Day1{}
	task.Init()
	tasks.RunTaskTimedAverage(task, 1000000)
}
