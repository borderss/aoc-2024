package tasks

import (
	"fmt"
	"os"
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

	// fragementedString = [[0        ] [. .       ] [1 1 1      ] [. . . .     ] [2 2 2 2 2    ]]
	// flattened = "0...111.....22222"

	// iterate over the fragments, move the last data blocks to the first available free space ('.'), like this:
	// 0..111....22222
	// 02.111....2222.
	// 022111....222..
	// 0221112...22...
	// 02211122..2....
	// 022111222......

	fmt.Println(fragmentedString)

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

	fmt.Println(fragmentedString)

	var flattened string
	for _, v := range fragmentedString {
		for _, vv := range v {
			if vv == "." || vv == "" {
				continue
			}
			flattened += vv
		}
	}

	var sum int64
	for i := 0; i < len(flattened); i++ {
		sum += int64(int64(i) * int64(flattened[i]-48))
	}

	return sum, nil
}

func move(s string, from, to int) string {
	runes := []rune(s)
	runes[to] = runes[from]
	runes[from] = '.'
	return string(runes)
}

func (d *Day9) Puzzle2() (any, error) {

	return 0, nil
}
