package search

import "golang.org/x/exp/constraints"

// Linear searches for an element in a slice of ordered or unordered elements,
// returning the index of the found element. The time complexity of the linear search is
// O(n) where n is the number of elements in the slice.
func Linear[V constraints.Ordered](s []V, x V) (int, bool) {
	for i, v := range s {
		if v == x {
			return i, true
		}
	}
	return 0, false
}
