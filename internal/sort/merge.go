package sort

import (
	"golang.org/x/exp/constraints"
)

// Merge sorts a slice of orderable elements by dividing the slice into smaller
// slices, sorting each slice, and then merging them back together (divide and
// conquer). The time complexity of the merge sort is O(n log n), as it is always
// dividing the slices in half for processing. The space complexity is O(n), as
// it is creating a new slice for each merge.
func Merge[V constraints.Ordered](s []V, asc bool) {
	copy(s, merge(s, asc))
}

func merge[V constraints.Ordered](s []V, asc bool) []V {
	if len(s) <= 1 {
		return s
	}

	m := int(len(s) / 2)
	l := merge(s[:m], asc)
	r := merge(s[m:], asc)

	rs := make([]V, len(s))
	i, j, k := 0, 0, 0
	for i < len(l) && j < len(r) {
		if asc {
			if l[i] < r[j] {
				rs[k] = l[i]
				i++
			} else {
				rs[k] = r[j]
				j++
			}
		} else {
			if l[i] > r[j] {
				rs[k] = l[i]
				i++
			} else {
				rs[k] = r[j]
				j++
			}
		}
		k++
	}

	for i < len(l) {
		rs[k] = l[i]
		i++
		k++
	}

	for j < len(r) {
		rs[k] = r[j]
		j++
		k++
	}

	return rs
}
