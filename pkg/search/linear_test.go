package search_test

import (
	"testing"

	"github.com/prichrd/concepts/pkg/search"
	"github.com/stretchr/testify/require"
)

func TestLinear(t *testing.T) {
	t.Run("leftmost_element", func(t *testing.T) {
		index, exists := search.Linear([]int{1, 2, 3, 4, 5}, 3)
		require.True(t, exists)
		require.Equal(t, 2, index)
	})
	t.Run("rightmost_element", func(t *testing.T) {
		index, exists := search.Linear([]int{1, 2, 3, 4, 5}, 5)
		require.True(t, exists)
		require.Equal(t, 4, index)
	})
	t.Run("middle_element", func(t *testing.T) {
		index, exists := search.Linear([]int{1, 2, 3, 4, 5}, 3)
		require.True(t, exists)
		require.Equal(t, 2, index)
	})
	t.Run("missing_element", func(t *testing.T) {
		_, exists := search.Linear([]int{1, 2, 3, 4, 5}, 999)
		require.False(t, exists)
	})
}
