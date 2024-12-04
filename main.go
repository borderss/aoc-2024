package main

import (
	"fmt"

	"github.com/borderss/aoc-2024/tasks"
)

func main() {
	task := &tasks.Day4{}

	fmt.Printf("running task: %T \n", *task)
	task.Init()
	tasks.RunTaskTimedAverage(task, 1000)
}
