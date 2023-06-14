package fuzzy

// Match represents a string match.
type Match struct {
	Hits     []int
	Distance int
}

// Matcher is implemented by types able to perform fuzzy string matching.
type Matcher interface {
	// Match matches a string (needle) in another string (haystack).
	Match(haystack, needle string) *Match
}
