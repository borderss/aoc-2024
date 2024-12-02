package tasks

import (
	"container/heap"
	"os"
	"strconv"
	s "strings"
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

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (d *Day1) Puzzle1() (any, error) {
	var heap1 IntHeap
	var heap2 IntHeap
	var sum int

	heap.Init(&heap1)
	heap.Init(&heap2)

	d.data = s.ReplaceAll(d.data, "\n", ",")
	d.data = s.ReplaceAll(d.data, "   ", ",")
	arr := s.Split(d.data, ",")
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			val, _ := strconv.Atoi(arr[i])
			heap.Push(&heap1, val)
		} else {
			val, _ := strconv.Atoi(arr[i])
			heap.Push(&heap2, val)
		}
	}

	for heap1.Len() > 0 {
		s := heap.Pop(&heap1).(int) - heap.Pop(&heap2).(int)
		if s < 0 {
			s = -s
		}
		sum += s
	}

	return sum, nil
}

func (d *Day1) Puzzle2() (any, error) {
	var arr1, arr2 []int

	d.data = s.ReplaceAll(d.data, "\n", ",")
	d.data = s.ReplaceAll(d.data, "   ", ",")
	arr := s.Split(d.data, ",")
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
