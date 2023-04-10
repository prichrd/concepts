package sort

import (
	"golang.org/x/exp/constraints"
)

// Quick sorts a slice of orderable elements by dividing the problem into smaller
// slices and sorting them recursively (divide and conquer). It uses a pivot element
// to partition the slice into two sub-slices. The elements smaller than the pivot
// are placed to the left of the pivot and the elements greater than the pivot are
// placed to the right of the pivot. The pivot element is then placed in its correct
// position in the sorted slice. The sub-slices are then sorted recursively.
// The time complexity of the quick sort algorithm is O(n log n) in the average case
// and O(n^2) in the worst case. The space complexity is O(log n) in the average case
// and O(n) in the worst case.
func Quick[V constraints.Ordered](s []V, asc bool) {
	quickSort(s, 0, len(s)-1, asc)
}

func quickSort[V constraints.Ordered](s []V, from, to int, asc bool) {
	if from >= to {
		return
	}

	p := partition(s, from, to, asc)
	quickSort(s, from, p-1, asc)
	quickSort(s, p+1, to, asc)
}

func partition[V constraints.Ordered](s []V, from, to int, asc bool) int {
	pivot := to
	l := from - 1
	for r := from; r < to; r++ {
		if asc && s[r] <= s[pivot] {
			l++
			s[r], s[l] = s[l], s[r]
		} else if !asc && s[r] >= s[pivot] {
			l++
			s[r], s[l] = s[l], s[r]
		}
	}
	s[pivot], s[l+1] = s[l+1], s[pivot]
	return l + 1
}
