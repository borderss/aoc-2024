package tasks

import (
	"os"
	"strconv"
	"strings"

	"github.com/borderss/aoc-2024/utils"
)

type Day11 struct {
	data string
}

func (d *Day11) Init() error {
	val, err := os.ReadFile("./data/day11.txt")
	if err != nil {
		return err
	}

	d.data = string(val)
	return nil
}

func (d *Day11) Puzzle1() (any, error) {
	values := strings.Split(d.data, " ")

	freq := make(map[string]int)
	for _, s := range values {
		freq[s]++
	}

	memo := make(map[string]map[string]int)

	finalFreq := make(map[string]int)
	for val, count := range freq {
		res := transform(val, 25, memo)
		for k, v := range res {
			finalFreq[k] += v * count
		}
	}

	totalCount := 0
	for _, c := range finalFreq {
		totalCount += c
	}
	return totalCount, nil
}

func (d *Day11) Puzzle2() (any, error) {
	values := strings.Split(d.data, " ")

	freq := make(map[string]int)
	for _, s := range values {
		freq[s]++
	}

	memo := make(map[string]map[string]int)

	finalFreq := make(map[string]int)
	for val, count := range freq {
		res := transform(val, 75, memo)
		for k, v := range res {
			finalFreq[k] += v * count
		}
	}

	totalCount := 0
	for _, c := range finalFreq {
		totalCount += c
	}
	return totalCount, nil
}

func transform(value string, steps int, memo map[string]map[string]int) map[string]int {
	if steps == 0 {
		return map[string]int{value: 1}
	}

	key := value + ":" + strconv.Itoa(steps)
	if cached, ok := memo[key]; ok {
		return cached
	}

	next := blink(value)
	result := make(map[string]int)
	for _, stone := range next {
		partial := transform(stone, steps-1, memo)
		for k, v := range partial {
			result[k] += v
		}
	}

	memo[key] = result
	return result
}

func blink(value string) []string {
	if value == "0" {
		return []string{"1"}
	}

	digits := len(value)
	if digits%2 == 0 {
		part1 := value[:digits/2]
		par2 := value[digits/2:]

		par2 = strings.TrimLeft(par2, "0")
		if par2 == "" {
			par2 = "0"
		}
		return []string{part1, par2}
	}

	value = strconv.FormatInt(utils.ParseInt[int64](value)*2024, 10)

	return []string{value}
}
