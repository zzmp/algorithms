package sort

import (
	"reflect"
	"testing"
)

type Sortable struct {
	Int   int
	Place int
}

func TestInsertion(t *testing.T) {
	less := func(a, b reflect.Value) bool {
		return a.Int() < b.Int()
	}

	list := []int{5, 2, 4, 6, 1, 3}
	sortedList := []int{1, 2, 3, 4, 5, 6}

	Insertion(list, less)

	for i := range list {
		if list[i] != sortedList[i] {
			t.Fatalf("%v != %v", list, sortedList)
		}
	}
}

func TestInsertionStable(t *testing.T) {
	less := func(a, b reflect.Value) bool {
		return a.FieldByName("Int").Int() < b.FieldByName("Int").Int()
	}

	list := []Sortable{{5, 0}, {2, 1}, {4, 3}, {2, 4}, {1, 5}, {3, 6}}
	sortedList := []Sortable{{1, 5}, {2, 1}, {2, 4}, {3, 6}, {4, 3}, {5, 0}}

	Insertion(list, less)

	for i := range list {
		if list[i] != sortedList[i] {
			t.Fatalf("%v != %v", list, sortedList)
		}
	}
}
