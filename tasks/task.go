package tasks

import (
	"fmt"
	"time"
)

type Task interface {
	Init() error
	Puzzle1() (any, error)
	Puzzle2() (any, error)
}

func RunTask(task Task) error {
	RunTaskPart1(task)
	RunTaskPart2(task)
	return nil
}

func RunTaskPart1(task Task) error {
	var result1 any

	result1, err := task.Puzzle1()
	if err != nil {
		return err
	}
	fmt.Println("result:", result1)

	return nil
}

func RunTaskPart2(task Task) error {
	var result2 any

	result2, err := task.Puzzle2()
	if err != nil {
		return err
	}
	fmt.Println("result:", result2)

	return nil
}

func RunTaskTimed(task Task) error {
	RunTaskTimedPart1(task)
	RunTaskTimedPart2(task)
	return nil
}

func RunTaskTimedPart1(task Task) error {
	var result1 any
	var err error

	t1 := time.Now()
	result1, err = task.Puzzle1()
	if err != nil {
		return err
	}
	puzzle1Time := time.Since(t1)
	fmt.Println("Puzzle 1 time:", puzzle1Time)
	fmt.Println("result:", result1)

	return nil
}

func RunTaskTimedPart2(task Task) error {
	var result2 any
	var err error

	t2 := time.Now()
	result2, err = task.Puzzle2()
	if err != nil {
		return err
	}
	puzzle2Time := time.Since(t2)
	fmt.Println("Puzzle 2 time:", puzzle2Time)
	fmt.Println("result:", result2)

	return nil
}

func RunTaskTimedAverage(task Task, times int) error {
	fmt.Println("Average run time over", times, "runs")
	fmt.Println()
	RunTaskTimedAveragePart1(task, times)
	RunTaskTimedAveragePart2(task, times)
	return nil
}

func RunTaskTimedAveragePart1(task Task, times int) error {
	var result1 any
	var err error

	t1 := time.Now()
	for i := 0; i < times; i++ {
		result1, err = task.Puzzle1()
		if err != nil {
			return err
		}
	}
	puzzle1Time := time.Since(t1) / time.Duration(times)
	fmt.Println("Puzzle 1 time:", puzzle1Time)
	fmt.Println("result:", result1)

	return nil
}

func RunTaskTimedAveragePart2(task Task, times int) error {
	var result2 any
	var err error

	t2 := time.Now()
	for i := 0; i < times; i++ {
		result2, err = task.Puzzle2()
		if err != nil {
			return err
		}
	}
	puzzle2Time := time.Since(t2) / time.Duration(times)
	fmt.Println("Puzzle 2 time:", puzzle2Time)
	fmt.Println("result:", result2)

	return nil
}
