// BenchmarkInsertionReflect-4   2000000   822 ns/op  112 B/op 6 allocs/op
// BenchmarkInsertionInterface-4 30000000  48.7 ns/op 0 B/op   0 allocs/op
// BenchmarkInsertionInt-4       100000000 12.0 ns/op 0 B/op   0 allocs/op
package sort

import (
	"reflect"
	"testing"
)

var (
	resultList []int
)

func BenchmarkInsertionReflect(b *testing.B) {
	less := func(a, b reflect.Value) bool {
		return a.Int() < b.Int()
	}

	list := getList()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		InsertionReflect(list, less)
	}

	// Store the result to avoid compiler optimization
	resultList = list
}

func BenchmarkInsertionInterface(b *testing.B) {
	list := IntList(getList())

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		InsertionInterface(&list)
	}

	// Store the result to avoid compiler optimization
	resultList = list
}

func BenchmarkInsertionInt(b *testing.B) {
	list := getList()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		InsertionInt(list)
	}

	// Store the result to avoid compiler optimization
	resultList = list
}
