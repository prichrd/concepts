package fuzzy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNaiveMatcher(t *testing.T) {
	matcher := &NaiveMatcher{}
	t.Run("matched_sparsed", func(t *testing.T) {
		m := matcher.Match("abcdefg", "aeg")
		require.Equal(t, m, &Match{
			Hits:     []int{0, 4, 6},
			Distance: 4,
		})
	})
	t.Run("matched_beg", func(t *testing.T) {
		m := matcher.Match("abcdefg", "abc")
		require.Equal(t, m, &Match{
			Hits:     []int{0, 1, 2},
			Distance: 4,
		})
	})
	t.Run("matched_med", func(t *testing.T) {
		m := matcher.Match("abcdefg", "bcd")
		require.Equal(t, m, &Match{
			Hits:     []int{1, 2, 3},
			Distance: 4,
		})
	})
	t.Run("matched_end", func(t *testing.T) {
		m := matcher.Match("abcdefg", "efg")
		require.Equal(t, m, &Match{
			Hits:     []int{4, 5, 6},
			Distance: 4,
		})
	})
	t.Run("matched_all", func(t *testing.T) {
		m := matcher.Match("abcdefg", "abcdefg")
		require.Equal(t, m, &Match{
			Hits:     []int{0, 1, 2, 3, 4, 5, 6},
			Distance: 0,
		})
	})
	t.Run("matched_same_letters", func(t *testing.T) {
		m := matcher.Match("abba", "bb")
		require.Equal(t, m, &Match{
			Hits:     []int{1, 2},
			Distance: 2,
		})
	})
	t.Run("matched_empty_needle", func(t *testing.T) {
		m := matcher.Match("abcdefg", "")
		require.Equal(t, m, &Match{
			Hits:     []int{},
			Distance: 7,
		})
	})
	t.Run("unmatched_smaller_haystack", func(t *testing.T) {
		m := matcher.Match("a", "abc")
		require.Nil(t, m)
	})
	t.Run("unmatched_not_present", func(t *testing.T) {
		m := matcher.Match("abcdefg", "xyz")
		require.Nil(t, m)
	})
	t.Run("unmatched_only_last_elem", func(t *testing.T) {
		m := matcher.Match("abcdefg", "efx")
		require.Nil(t, m)
	})
}
