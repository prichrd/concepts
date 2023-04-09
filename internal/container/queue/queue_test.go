package queue_test

import (
	"testing"

	"github.com/prichrd/concepts/internal/container/queue"
	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	q := queue.NewQueue[int]()
	require.Nil(t, q.Pop())
	require.Equal(t, 0, q.Len())
	q.Push(10)
	require.Equal(t, 1, q.Len())
	q.Push(20)
	require.Equal(t, 2, q.Len())
	v := q.Peek()
	require.Equal(t, 10, *v)
	v = q.Pop()
	require.Equal(t, 10, *v)
	require.Equal(t, 1, q.Len())
	v = q.Pop()
	require.Equal(t, 20, *v)
	require.Equal(t, 0, q.Len())
	require.Nil(t, q.Pop())
}
