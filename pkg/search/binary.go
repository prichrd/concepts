package search

import (
	"golang.org/x/exp/constraints"
)

// Binary searches for an element in a slice of ordered elements, returning the index
// of the found element. The time complexity of the binary search is O(log n), where n
// is the number of elements in the slice.
func Binary[V constraints.Ordered](s []V, x V) (int, bool) {
	l, r := 0, len(s)-1
	for l <= r {
		mid := l + (r-l)/2
		if s[mid] == x {
			return mid, true
		} else if s[mid] < x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return 0, false
}
