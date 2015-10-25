package reflection

import "reflect"

type Less func(i, j reflect.Value) bool
type Sortable interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// Insertion sort.
//
// Relies on reflection to sort an interface{} given a Less func.
//
// Panics if list is not Array, Chan, Map, or Slice.
func InsertionReflect(list interface{}, less Less) {
	value := reflect.ValueOf(list)

	// Panics if Type is not Array, Chan, Map, Ptr, or Slice.
	elemType := reflect.TypeOf(list).Elem()

	// value.Len() panics if Kind is not Array, Chan, Map, Slice, or String.
	for j := 1; j < value.Len(); j++ {
		// Store list[j] in a new variable
		key := reflect.Indirect(reflect.New(elemType))
		key.Set(value.Index(j))

		// Insert list[j] into the sorted sequence list[0,j-1]
		i := j - 1
		for i >= 0 && less(key, value.Index(i)) {
			value.Index(i + 1).Set(value.Index(i))
			i--
		}
		value.Index(i + 1).Set(key)
	}
}

// Insertion sort.
//
// Relies on Sortable interface methods.
//
// Suffers some performance by only having a Swap operation,
// as insertion sort usually stores a key in memory to have
// at most j assignments per loop, where j is the key index.
//
// Instead, InsertionInterface can have at most j*2 assignments,
// as the key is never stored externally
// (i.e. this is not a true insertion sort).
func InsertionInterface(list Sortable) {
	for j := 1; j < list.Len(); j++ {
		i := j - 1
		for i >= 0 && list.Less(i+1, i) {
			list.Swap(i+1, i)
			i--
		}
	}
}

// Insertion sort.
//
// Relies on built-in properties of []int.
func InsertionInt(list []int) {
	for j := 1; j < len(list); j++ {
		key := list[j]

		i := j - 1
		for i >= 0 && key < list[i] {
			list[i+1] = list[i]
			i--
		}
		list[i+1] = key
	}
}
