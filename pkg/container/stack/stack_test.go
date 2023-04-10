package stack_test

import (
	"testing"

	"github.com/prichrd/concepts/pkg/container/stack"
	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	s := stack.NewStack[int]()
	require.Nil(t, s.Pop())
	require.Equal(t, 0, s.Len())
	s.Push(10)
	require.Equal(t, 1, s.Len())
	s.Push(20)
	require.Equal(t, 2, s.Len())
	v := s.Peek()
	require.Equal(t, 20, *v)
	v = s.Pop()
	require.Equal(t, 20, *v)
	require.Equal(t, 1, s.Len())
	v = s.Pop()
	require.Equal(t, 10, *v)
	require.Equal(t, 0, s.Len())
	require.Nil(t, s.Pop())
}
