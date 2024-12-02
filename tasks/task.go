package tasks

import (
	"fmt"
	"time"
)

type Task interface {
	Init() error
	Puzzle1() (string, error)
	Puzzle2() (string, error)
}

func RunTaskTimed(task Task) error {
	err := task.Init()
	if err != nil {
		return err
	}

	var result1, result2 string

	t1 := time.Now()
	result1, err = task.Puzzle1()
	if err != nil {
		return err
	}
	puzzle1Time := time.Since(t1)
	fmt.Println("Puzzle 1 time:", puzzle1Time)
	fmt.Println("result:", result1)
	fmt.Println("")

	// -------

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

func RunTask(task Task) error {
	err := task.Init()
	if err != nil {
		return err
	}

	var result1, result2 string

	result1, err = task.Puzzle1()
	if err != nil {
		return err
	}
	fmt.Println("result:", result1)
	fmt.Println("")

	// -------

	result2, err = task.Puzzle2()
	if err != nil {
		return err
	}
	fmt.Println("result:", result2)

	return nil
}

func RunTaskTimedAverage(task Task, times int) error {
	err := task.Init()
	if err != nil {
		return err
	}

	var result1, result2 string

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
	fmt.Println("")

	// -------

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
