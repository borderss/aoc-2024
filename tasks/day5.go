package tasks

import (
	"fmt"
	"os"
	"strings"

	"github.com/borderss/aoc-2024/utils"
)

type Day5 struct {
	data string
}

func (d *Day5) Init() error {
	val, err := os.ReadFile("./data/day5.txt")
	if err != nil {
		return err
	}

	d.data = string(val)
	return nil
}

func (d *Day5) Puzzle1() (any, error) {
	sections := strings.Split(d.data, "\n\n")
	rulesRaw := strings.Split(sections[0], "\n")
	updatesRaw := strings.Split(sections[1], "\n")

	rules := make([][]int8, 0, len(rulesRaw))
	for _, v := range rulesRaw {
		rule := make([]int8, 2)
		fmt.Sscanf(v, "%d|%d", &rule[0], &rule[1])
		rules = append(rules, rule)
	}

	updates := make([][]int8, 0, len(updatesRaw))
	for _, v := range updatesRaw {
		values := strings.Split(v, ",")
		castRow := make([]int8, 0, len(values))
		for _, v := range values {
			castRow = append(castRow, utils.ParseInt[int8](v))
		}

		updates = append(updates, castRow)
	}

	sum := 0

	for _, update := range updates {
		for i, rule := range rules {
			passes := rulePasses(rule, update)
			if passes == -1 {
				break
			}

			if i == len(rules)-1 {
				sum += int(update[len(update)/2])
				break
			}
		}
	}

	return sum, nil
}

func (d *Day5) Puzzle2() (any, error) {
	sections := strings.Split(d.data, "\n\n")
	rulesRaw := strings.Split(sections[0], "\n")
	updatesRaw := strings.Split(sections[1], "\n")

	rules := make([][]int8, 0, len(rulesRaw))
	for _, v := range rulesRaw {
		rule := make([]int8, 2)
		fmt.Sscanf(v, "%d|%d", &rule[0], &rule[1])
		rules = append(rules, rule)
	}

	updates := make([][]int8, 0, len(updatesRaw))
	for _, v := range updatesRaw {
		values := strings.Split(v, ",")
		castRow := make([]int8, 0, len(values))
		for _, v := range values {
			castRow = append(castRow, utils.ParseInt[int8](v))
		}

		updates = append(updates, castRow)
	}

	var incorrectUpdates [][]int8

	for _, update := range updates {
		for _, rule := range rules {
			passes := rulePasses(rule, update)
			if passes == -1 {
				incorrectUpdates = append(incorrectUpdates, update)
				break
			}
		}
	}

	fixedUpdates := make([][]int8, 0, len(incorrectUpdates))
	for _, update := range incorrectUpdates {
		for !ruleArrayPasses(rules, update) {
			for i, rule := range rules {
				for rulePasses(rule, update) == -1 {
					update = ruleSwap(rule, update)
				}
				if i == len(rules)-1 {
					fixedUpdates = append(fixedUpdates, update)
					break
				}
			}
		}
	}

	fixedUpdates = removeDuplicateSlices(fixedUpdates)

	sum := 0

	for _, update := range fixedUpdates {
		sum += int(update[len(update)/2])
	}

	return sum, nil
}

func ruleSwap(rule []int8, update []int8) []int8 {
	if hasElement(update, rule[0]) && hasElement(update, rule[1]) {
		if indexOf(update, rule[0]) < indexOf(update, rule[1]) {
			return update
		}

		temp := update[indexOf(update, rule[0])]
		update[indexOf(update, rule[0])] = update[indexOf(update, rule[1])]
		update[indexOf(update, rule[1])] = temp

		return update
	}

	return update
}

func rulePasses(rule []int8, update []int8) int {
	if hasElement(update, rule[0]) && hasElement(update, rule[1]) {
		if indexOf(update, rule[0]) < indexOf(update, rule[1]) {
			return 1
		}

		return -1
	}

	return 0
}

func ruleArrayPasses(rule [][]int8, update []int8) bool {
	for _, r := range rule {
		if rulePasses(r, update) == -1 {
			return false
		}
	}

	return true
}

func hasElement(slice []int8, value int8) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}

func removeDuplicateSlices(slice [][]int8) [][]int8 {
	keys := make(map[string]bool)
	list := [][]int8{}

	for _, entry := range slice {
		if _, value := keys[fmt.Sprint(entry)]; !value {
			keys[fmt.Sprint(entry)] = true
			list = append(list, entry)
		}
	}

	return list
}

func indexOf(slice []int8, value any) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}

	return -1
}
