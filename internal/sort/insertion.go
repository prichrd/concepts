package sort

import "golang.org/x/exp/constraints"

// Insertion sorts a slice of orderable elements by picking an element and
// inserting it into the correct position in the sorted slice. The time
// complexity of the insertion sort is O(n^2), as it compares each element
// with every other element in the slice. The space complexity is O(1), as
// the sort is done in place.
func Insertion[V constraints.Ordered](s []V, asc bool) {
	for i := 1; i < len(s); i++ {
		curr := s[i]
		j := i - 1
		for asc && j >= 0 && curr < s[j] {
			s[j+1] = s[j]
			j--
		}
		for !asc && j >= 0 && curr > s[j] {
			s[j+1] = s[j]
			j--
		}
		s[j+1] = curr
	}
}
