package lru_test

import (
	"testing"

	"github.com/prichrd/concepts/internal/cache/lru"
	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("puts_and_gets", func(t *testing.T) {
		c := lru.New[string, int](5)
		c.Put("a", 1)                    // a
		c.Put("b", 2)                    // b a
		c.Put("c", 3)                    // c b a
		c.Put("d", 4)                    // d c b a
		c.Put("e", 5)                    // e d c b a
		require.Equal(t, 1, *c.Get("a")) // a e d c b
		c.Put("f", 6)                    // f a e d c
		require.Nil(t, c.Get("b"))
		require.Equal(t, 1, *c.Get("a"))
		require.Equal(t, 3, *c.Get("c"))
		require.Equal(t, 4, *c.Get("d"))
		require.Equal(t, 5, *c.Get("e"))
		require.Equal(t, 6, *c.Get("f"))
	})

	t.Run("updates", func(t *testing.T) {
		c := lru.New[string, int](2)
		c.Put("a", 1) // a
		c.Put("b", 2) // b a
		c.Put("a", 3) // a b
		c.Put("c", 4) // c a
		require.Equal(t, 3, *c.Get("a"))
		require.Equal(t, 4, *c.Get("c"))
	})
}
