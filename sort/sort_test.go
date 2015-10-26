package sort

import (
	"sort"
	"testing"
)

func test(t *testing.T, algorithm func([]int)) {
	list := []int{5, 3, 4, 1, 7, 2, 3, 8, 5, 3, 8, 9, 0, 6, 4}
	algorithm(list)

	if !sort.IntsAreSorted(list) {
		sortedList := make([]int, len(list))
		copy(sortedList, list)
		sort.Ints(sortedList)

		t.Fatalf("%v != %v ", list, sortedList)
	}
}

// Insertion
func TestInsertion(t *testing.T) { test(t, Insertion) }

// Merge
func TestMerge(t *testing.T) { test(t, Merge) }
