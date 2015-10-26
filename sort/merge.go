package sort

func merge(list []int, lo, mid, hi int) {
	L := make([]int, mid-lo)
	R := make([]int, hi-mid)
	copy(L, list[lo:mid])
	copy(R, list[mid:hi])

	var l, r int
	for i := lo; i < hi; i++ {
		if l == len(L) {
			list[i] = R[r]
			r++
		} else if r == len(R) {
			list[i] = L[l]
			l++
		} else if L[l] <= R[r] {
			list[i] = L[l]
			l++
		} else {
			list[i] = R[r]
			r++
		}
	}
}

func mergeSort(list []int, lo, hi int) {
	if lo < hi-1 {
		mid := (lo + hi + 1) / 2
		mergeSort(list, lo, mid)
		mergeSort(list, mid, hi)
		merge(list, lo, mid, hi)
	}
}

func Merge(list []int) {
	mergeSort(list, 0, len(list))
}
