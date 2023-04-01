package sort

import "golang.org/x/exp/constraints"

// Bubble sorts a slice of orderable elements by bubbling the smallest
// or largest element to the end of a slice. The time complexity of the
// bubble sort is O(n^2), as it requires two nested loops. The space complexity
// is O(1), as the sort is done in place.
func Bubble[V constraints.Ordered](s []V, asc bool) {
	for step := range s {
		for i := 0; i < len(s)-step-1; i++ {
			if asc && s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
			} else if !asc && s[i] < s[i+1] {
				s[i+1], s[i] = s[i], s[i+1]
			}
		}
	}
}
