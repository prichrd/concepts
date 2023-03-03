package list_test

import (
	"testing"

	"github.com/prichrd/concepts/internal/container/list"
	"github.com/stretchr/testify/require"
)

func TestDoublyLinkedList(t *testing.T) {
	t.Run("push_front_and_back", func(t *testing.T) {
		ll := list.NewList[int]()
		ll.PushFront(1)
		ll.PushBack(2)
		ll.PushBack(3)
		ll.PushFront(0)
		assetDoublyLinkedList(t, ll, []int{0, 1, 2, 3})
	})

	t.Run("remove_front_element", func(t *testing.T) {
		ll := list.NewList[int]()
		ll.PushBack(1)
		ll.PushBack(2)
		ll.PushBack(3)
		ll.Remove(ll.Front())
		assetDoublyLinkedList(t, ll, []int{2, 3})
	})

	t.Run("remove_mid_element", func(t *testing.T) {
		ll := list.NewList[int]()
		ll.PushBack(1)
		ll.PushBack(2)
		ll.PushBack(3)
		ll.Remove(ll.Front().Next())
		assetDoublyLinkedList(t, ll, []int{1, 3})
	})

	t.Run("remove_back_element", func(t *testing.T) {
		ll := list.NewList[int]()
		ll.PushBack(1)
		ll.PushBack(2)
		ll.PushBack(3)
		ll.Remove(ll.Back())
		assetDoublyLinkedList(t, ll, []int{1, 2})
	})

	t.Run("move_front_to_front", func(t *testing.T) {
		ll := list.NewList[int]()
		ll.PushBack(1)
		ll.PushBack(2)
		ll.PushBack(3)
		ll.MoveToFront(ll.Front())
		assetDoublyLinkedList(t, ll, []int{1, 2, 3})
	})

	t.Run("move_mid_to_front", func(t *testing.T) {
		ll := list.NewList[int]()
		ll.PushBack(1)
		ll.PushBack(2)
		ll.PushBack(3)
		ll.MoveToFront(ll.Front().Next())
		assetDoublyLinkedList(t, ll, []int{2, 1, 3})
	})

	t.Run("move_back_to_front", func(t *testing.T) {
		ll := list.NewList[int]()
		ll.PushBack(1)
		ll.PushBack(2)
		ll.PushBack(3)
		ll.MoveToFront(ll.Back())
		assetDoublyLinkedList(t, ll, []int{3, 1, 2})
	})

	t.Run("empty_list_getters", func(t *testing.T) {
		ll := list.NewList[int]()
		require.Nil(t, ll.Front())
		require.Nil(t, ll.Back())
		require.Zero(t, ll.Len())
	})
}

func assetDoublyLinkedList[V any](t *testing.T, ll *list.List[V], exp []V) {
	forward := make([]V, 0)
	n := ll.Front()
	for {
		forward = append(forward, n.Val())
		if n.Next() == nil {
			break
		}
		n = n.Next()
	}

	backward := make([]V, 0)
	n = ll.Back()
	for {
		backward = append(backward, n.Val())
		if n.Prev() == nil {
			break
		}
		n = n.Prev()
	}

	rbackward := make([]V, len(backward))
	for i := 0; i < len(rbackward); i++ {
		rbackward[i] = backward[len(backward)-i-1]
	}

	require.Equal(t, exp, forward, "forward list doesn't match expectations")
	require.Equal(t, exp, rbackward, "backward list doesn't match expectations")
	require.Equal(t, len(exp), ll.Len(), "the list is of wrong length")
}
