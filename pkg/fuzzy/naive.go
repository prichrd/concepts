package fuzzy

// NaiveMatcher implements a naive fuzzy string matching algorithm. It exposes
// a Match method that returns a Match object if the needle is found in the
// haystack, or nil otherwise.
type NaiveMatcher struct{}

// Match matches a string (needle) in another string (haystack).
func (m *NaiveMatcher) Match(haystack, needle string) *Match {
	if len(haystack) < len(needle) {
		return nil
	}

	r := &Match{
		Hits: make([]int, len(needle)),
	}

	n := 0
	for h := 0; h < len(haystack); h++ {
		if n == len(needle) {
			break
		}
		if haystack[h] == needle[n] {
			r.Hits[n] = h
			n++
		}
	}

	if n != len(needle) {
		return nil
	}

	r.Distance = len(haystack) - len(needle)

	return r
}
