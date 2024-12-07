package tasks

import (
	"os"
	"strconv"
	"strings"

	"slices"

	"github.com/borderss/aoc-2024/utils"
)

type Day1 struct {
	data string
}

func (d *Day1) Init() error {
	val, err := os.ReadFile("./data/day1.txt")
	if err != nil {
		return err
	}

	d.data = string(val)
	return nil
}

func (d *Day1) Puzzle1() (any, error) {
	data := strings.ReplaceAll(d.data, "\n", ",")
	data = strings.ReplaceAll(data, "   ", ",")
	arr := strings.Split(data, ",")

	even := make([]int32, 0, 1000)
	odd := make([]int32, 0, 1000)

	for i, v := range arr {
		val := utils.ParseInt[int32](v)
		if i%2 == 0 {
			even = append(even, val)
		} else {
			odd = append(odd, val)
		}
	}

	slices.Sort(even)
	slices.Sort(odd)

	var sum int32
	for i := range even {
		s := even[i] - odd[i]
		if s < 0 {
			s = -s
		}
		sum += s
	}

	return sum, nil
}

func (d *Day1) Puzzle2() (any, error) {
	var arr1, arr2 []int

	d.data = strings.ReplaceAll(d.data, "\n", ",")
	d.data = strings.ReplaceAll(d.data, "   ", ",")
	arr := strings.Split(d.data, ",")
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			val, _ := strconv.Atoi(arr[i])
			arr1 = append(arr1, val)
		} else {
			val, _ := strconv.Atoi(arr[i])
			arr2 = append(arr2, val)
		}
	}

	countMap := make(map[int]int)
	for i := 0; i < len(arr1); i++ {
		countMap[arr1[i]]++
	}

	sum := 0
	for i := 0; i < len(arr2); i++ {
		if countMap[arr2[i]] > 0 {
			sum += arr2[i] * countMap[arr2[i]]
		}
	}

	return sum, nil
}
