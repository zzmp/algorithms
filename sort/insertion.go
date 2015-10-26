package sort

func Insertion(list []int) {
	// Loop invariant: list[0..j] is sorted
	// Initialization: list[0..1] is trivially sorted
	// Termination: j == list.Len() => list is sorted
	for j := 1; j < len(list); j++ {
		key := list[j]

		// Maintenance: move list[j] until list[0..j] is sorted
		i := j
		for ; i > 0 && key < list[i-1]; i-- {
			list[i] = list[i-1]
		}
		list[i] = key
	}
}
