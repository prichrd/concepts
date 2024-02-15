package bst_test

import (
	"fmt"
	"testing"

	"github.com/prichrd/concepts/pkg/container/bst"
	"github.com/stretchr/testify/require"
)

func TestBinarySearchTree(t *testing.T) {
	t.Run("remove_empty_tree", func(t *testing.T) {
		bst := bst.New[int, string]()
		v, ok := bst.Remove(0)
		require.False(t, ok)
		require.Equal(t, "", v)
	})
	t.Run("remove_root_node", func(t *testing.T) {
		bst := bst.New[int, string]()
		bst.Insert(8, "8")
		bst.Insert(3, "3")
		bst.Insert(10, "10")
		bst.Insert(1, "1")
		bst.Insert(6, "6")
		bst.Insert(9, "9")
		v, ok := bst.Remove(8)
		require.True(t, ok)
		require.Equal(t, "8", v)
		elems := []int{1, 3, 6, 9, 10}
		i := 0
		bst.Walk(func(k int, v string) {
			require.Equal(t, elems[i], k)
			require.Equal(t, fmt.Sprintf("%d", elems[i]), v)
			i++
		})
	})
	t.Run("remove_subtree_root", func(t *testing.T) {
		bst := bst.New[int, string]()
		bst.Insert(8, "8")
		bst.Insert(3, "3")
		bst.Insert(10, "10")
		bst.Insert(1, "1")
		bst.Insert(6, "6")
		bst.Insert(9, "9")
		v, ok := bst.Remove(3)
		require.True(t, ok)
		require.Equal(t, "3", v)
		elems := []int{1, 6, 8, 9, 10}
		i := 0
		bst.Walk(func(k int, v string) {
			require.Equal(t, elems[i], k)
			require.Equal(t, fmt.Sprintf("%d", elems[i]), v)
			i++
		})
	})
	t.Run("remove_leaf", func(t *testing.T) {
		bst := bst.New[int, string]()
		bst.Insert(8, "8")
		bst.Insert(3, "3")
		bst.Insert(10, "10")
		bst.Insert(1, "1")
		bst.Insert(6, "6")
		bst.Insert(9, "9")
		v, ok := bst.Remove(1)
		require.True(t, ok)
		require.Equal(t, "1", v)
		elems := []int{3, 6, 8, 9, 10}
		i := 0
		bst.Walk(func(k int, v string) {
			require.Equal(t, elems[i], k)
			require.Equal(t, fmt.Sprintf("%d", elems[i]), v)
			i++
		})
	})

	t.Run("find_empty_tree", func(t *testing.T) {
		bst := bst.New[int, string]()
		v, ok := bst.Remove(0)
		require.False(t, ok)
		require.Equal(t, "", v)
	})
	t.Run("find_non_existant_node", func(t *testing.T) {
		bst := bst.New[int, string]()
		bst.Insert(8, "8")
		bst.Insert(3, "3")
		bst.Insert(10, "10")
		v, ok := bst.Find(99)
		require.False(t, ok)
		require.Equal(t, "", v)
	})
	t.Run("find_all_nodes", func(t *testing.T) {
		bst := bst.New[int, string]()
		bst.Insert(8, "8")
		bst.Insert(3, "3")
		bst.Insert(10, "10")
		bst.Insert(1, "1")
		v, ok := bst.Find(8)
		require.True(t, ok)
		require.Equal(t, "8", v)
		v, ok = bst.Find(3)
		require.True(t, ok)
		require.Equal(t, "3", v)
		v, ok = bst.Find(10)
		require.True(t, ok)
		require.Equal(t, "10", v)
		v, ok = bst.Find(1)
		require.True(t, ok)
		require.Equal(t, "1", v)
	})
	t.Run("walk_order", func(t *testing.T) {
		bst := bst.New[int, string]()
		bst.Insert(8, "8")
		bst.Insert(3, "3")
		bst.Insert(10, "10")
		bst.Insert(1, "1")
		bst.Insert(6, "6")
		bst.Insert(14, "14")
		bst.Insert(4, "4")
		bst.Insert(7, "7")
		bst.Insert(13, "13")
		elems := []int{1, 3, 4, 6, 7, 8, 10, 13, 14}
		i := 0
		bst.Walk(func(k int, v string) {
			require.Equal(t, elems[i], k)
			require.Equal(t, fmt.Sprintf("%d", elems[i]), v)
			i++
		})
	})
}
