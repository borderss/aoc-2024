package tasks

import (
	"fmt"
	"os"

	"github.com/borderss/aoc-2024/utils"
)

type Day9 struct {
	data string
}

func (d *Day9) Init() error {
	val, err := os.ReadFile("./data/day9.txt")
	if err != nil {
		return err
	}

	d.data = string(val)
	return nil
}

func (d *Day9) Puzzle1() (any, error) {
	fragmentedString := make([][]string, len(d.data))

	for i := 0; i < len(d.data); i++ {
		if i%2 == 0 {
			values := make([]string, int(d.data[i]-48))
			for j := 0; j < int(d.data[i]-48); j++ {
				values[j] = fmt.Sprintf("%d", i/2)
			}
			fragmentedString[i] = values
		} else {
			values := make([]string, int(d.data[i]-48))
			for j := 0; j < int(d.data[i]-48); j++ {
				values[j] = "."
			}
			fragmentedString[i] = values
		}
	}

	tempFragmentedString := fragmentedString

	for i := len(tempFragmentedString) - 1; i >= 0; i-- {
		for j := len(tempFragmentedString[i]) - 1; j >= 0; j-- {
			if tempFragmentedString[i][j] == "" || tempFragmentedString[i][j] == "." {
				continue
			}

			found := false
			for k := 0; k < len(fragmentedString) && !found; k++ {
				for l := 0; l < len(fragmentedString[k]); l++ {
					if fragmentedString[k][l] == "." {
						fragmentedString[k][l] = tempFragmentedString[i][j]
						tempFragmentedString[i][j] = "."
						found = true
						break
					}
				}
			}
		}
	}

	var flattenedArray []string
	for i := 0; i < len(fragmentedString); i++ {
		for j := 0; j < len(fragmentedString[i]); j++ {
			if fragmentedString[i][j] == "." {
				continue
			}

			flattenedArray = append(flattenedArray, fragmentedString[i][j])
		}
	}

	var sum int64
	for i := 0; i < len(flattenedArray); i++ {
		sum += int64(i) * utils.ParseInt[int64](flattenedArray[i])
	}

	return sum, nil
}

func (d *Day9) Puzzle2() (any, error) {
	fragmentedString := make([][]string, len(d.data))

	for i := 0; i < len(d.data); i++ {
		if i%2 == 0 {
			values := make([]string, int(d.data[i]-48))
			for j := 0; j < int(d.data[i]-48); j++ {
				values[j] = fmt.Sprintf("%d", i/2)
			}
			fragmentedString[i] = values
		} else {
			values := make([]string, int(d.data[i]-48))
			for j := 0; j < int(d.data[i]-48); j++ {
				values[j] = "."
			}
			fragmentedString[i] = values
		}
	}

	for i := len(fragmentedString) - 1; i >= 0; i-- {
		for j := 0; j < len(fragmentedString)-(len(fragmentedString)-i); j++ {
			if (len(fragmentedString[i]) == 0 || fragmentedString[i][0] == ".") ||
				(len(fragmentedString[j]) == 0 || fragmentedString[j][0] != ".") {
				continue
			}

			if len(fragmentedString[j]) >= len(fragmentedString[i]) {
				firstPart := make([]string, len(fragmentedString[i]))
				copy(firstPart, fragmentedString[i])
				secondPart := fragmentedString[j][len(fragmentedString[i]):]

				for k := range fragmentedString[i] {
					fragmentedString[i][k] = "."
				}

				fragmentedString = append(fragmentedString[:j], fragmentedString[j+1:]...)
				if len(secondPart) > 0 {
					fragmentedString = append(fragmentedString[:j], append([][]string{firstPart, secondPart}, fragmentedString[j:]...)...)
				} else {
					fragmentedString = append(fragmentedString[:j], append([][]string{firstPart}, fragmentedString[j:]...)...)
				}
			}
		}
	}

	var flattenedArray []string
	for i := 0; i < len(fragmentedString); i++ {
		for j := 0; j < len(fragmentedString[i]); j++ {
			flattenedArray = append(flattenedArray, fragmentedString[i][j])
		}
	}

	var sum int64
	for i := 0; i < len(flattenedArray); i++ {
		if flattenedArray[i] == "." {
			continue
		}

		sum += int64(i) * utils.ParseInt[int64](flattenedArray[i])
	}

	return sum, nil
}
