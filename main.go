package main

import (
	"fmt"

	"github.com/borderss/aoc-2024/tasks"
)

func main() {
	task := &tasks.Day1{}

	fmt.Printf("running task: %T \n", *task)
	task.Init()
	// tasks.RunTaskPart2(task)
	tasks.RunTaskTimedAverage(task, 1000)
}
