package lru_test

import (
	"testing"

	"github.com/prichrd/concepts/pkg/cache/lru"
	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("puts_and_gets", func(t *testing.T) {
		c := lru.New[string, int](5)
		c.Put("a", 1)                       // a
		c.Put("b", 2)                       // b a
		c.Put("c", 3)                       // c b a
		c.Put("d", 4)                       // d c b a
		c.Put("e", 5)                       // e d c b a
		expectKeyEquals(t, 1, true, "a", c) // a e d c b
		c.Put("f", 6)                       // f a e d c
		expectKeyEquals(t, 0, false, "b", c)
		expectKeyEquals(t, 1, true, "a", c)
		expectKeyEquals(t, 3, true, "c", c)
		expectKeyEquals(t, 4, true, "d", c)
		expectKeyEquals(t, 5, true, "e", c)
		expectKeyEquals(t, 6, true, "f", c)
	})

	t.Run("updates", func(t *testing.T) {
		c := lru.New[string, int](2)
		c.Put("a", 1) // a
		c.Put("b", 2) // b a
		c.Put("a", 3) // a b
		c.Put("c", 4) // c a
		expectKeyEquals(t, 3, true, "a", c)
		expectKeyEquals(t, 4, true, "c", c)
	})
}

func expectKeyEquals(t *testing.T, exp int, expOk bool, key string, c *lru.Cache[string, int]) {
	v, ok := c.Get(key)
	if !expOk {
		require.False(t, ok)
		return
	}
	require.Equal(t, exp, v)
}
