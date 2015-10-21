package sort

import (
	"fmt"
)

var (
	unsortedList []int = []int{5, 2, 4, 6, 1, 3}
	sortedList   []int = []int{1, 2, 3, 4, 5, 6}
)

func getList() []int {
	list := make([]int, len(unsortedList))
	copy(list, unsortedList)
	return list
}

func testList(list []int) error {
	for i := range sortedList {
		if list[i] != sortedList[i] {
			return fmt.Errorf("%v != %v", list, sortedList)
		}
	}

	return nil
}

type IntList []int

func (l *IntList) Len() int {
	return len([]int(*l))
}

func (l *IntList) Less(i, j int) bool {
	return []int(*l)[i] < []int(*l)[j]
}

func (l *IntList) Swap(i, j int) {
	t := []int(*l)[i]
	[]int(*l)[i] = []int(*l)[j]
	[]int(*l)[j] = t
}
