package sort

import (
	"golang.org/x/exp/constraints"
)

// Count sorts a slice of integer elements by counting the number of elements
// and positioning them in the correct order. The time complexity is O(n + k)
// where n is the number of elements and k is the maximum value. The space
// complexity is O(k) as we need to create a slice of length k to store the count
// of each elements.
func Count[V constraints.Integer](s []V, asc bool) {
	out := make([]V, len(s))

	if len(s) == 0 {
		return
	}

	max := s[0]
	for i := 1; i < len(s); i++ {
		if max < s[i] {
			max = s[i]
		}
	}

	acc := make([]uint, int(max)+1)
	for i := 0; i < len(s); i++ {
		acc[s[i]]++
	}

	for i := 1; i <= int(max); i++ {
		acc[i] += acc[i-1]
	}

	i := len(s) - 1
	for i >= 0 {
		if !asc {
			out[uint(len(s))-acc[s[i]]] = s[i]
			acc[s[i]] -= 1
		} else {
			out[acc[s[i]]-1] = s[i]
			acc[s[i]] -= 1
		}
		i -= 1
	}

	copy(s, out)
}
