package sort

import (
	"golang.org/x/exp/constraints"
)

// Selection sorts a slice of comparable elements by repeatedly finding the
// smallest or largest element and moving it to the beginning or end of the
// slice. The time complexity of the selection sort is O(n^2), as it makes
// n-1 passes through the slice and performs n-1 comparisons per pass.
func Selection[V constraints.Ordered](s []V, asc bool) {
	if len(s) == 0 {
		return
	}
	for step, _ := range s {
		l := step
		r := l
		mval := s[r]
		for i := l + 1; i < len(s); i++ {
			if asc && s[i] < mval {
				r = i
				mval = s[i]
			}
			if !asc && s[i] > mval {
				r = i
				mval = s[i]
			}
		}
		s[l], s[r] = s[r], s[l]
	}
}
