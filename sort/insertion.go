package sort

import "reflect"

type Less func(i, j reflect.Value) bool

// An insertion sort.
//
// A stable, in-place insertion sort.
// Panics if list is not Array, Chan, Map, or Slice.
//
func Insertion(list interface{}, less Less) {
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
