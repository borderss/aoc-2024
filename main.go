package main

import (
	"fmt"

	"github.com/borderss/aoc-2024/tasks"
)

func main() {
	task := &tasks.Day6{}

	fmt.Printf("running task: %T \n", *task)
	task.Init()
	tasks.RunTaskPart1(task)
	// tasks.RunTaskTimedAverage(task, 1000)
}
