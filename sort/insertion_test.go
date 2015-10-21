package sort

import (
	"reflect"
	"testing"
)

func TestInsertionReflect(t *testing.T) {
	less := func(a, b reflect.Value) bool {
		return a.Int() < b.Int()
	}

	list := getList()

	InsertionReflect(list, less)

	if err := testList(list); err != nil {
		t.Fatal(err)
	}
}

func TestInsertionReflectStable(t *testing.T) {
	type Stable struct {
		Int   int
		Place int
	}

	less := func(a, b reflect.Value) bool {
		return a.FieldByName("Int").Int() < b.FieldByName("Int").Int()
	}

	list := []Stable{{5, 0}, {2, 1}, {4, 3}, {2, 4}, {1, 5}, {3, 6}}
	sortedList := []Stable{{1, 5}, {2, 1}, {2, 4}, {3, 6}, {4, 3}, {5, 0}}

	InsertionReflect(list, less)

	for i := range sortedList {
		if list[i] != sortedList[i] {
			t.Fatalf("%v != %v", list, sortedList)
		}
	}
}

func TestInsertionInterface(t *testing.T) {
	list := IntList(getList())

	InsertionInterface(&list)

	if err := testList([]int(list)); err != nil {
		t.Fatal(err)
	}
}

func TestInsertionInt(t *testing.T) {
	list := getList()

	InsertionInt(list)

	if err := testList(list); err != nil {
		t.Fatal(err)
	}
}
