package main

import (
	"fmt"

	"github.com/borderss/aoc-2024/tasks"
)

func main() {
	task := &tasks.Day2{}

	fmt.Printf("running task: %T \n", *task)
	task.Init()
	// tasks.RunTaskPart2(task)
	tasks.RunTaskTimedAveragePart2(task, 1000)
}
