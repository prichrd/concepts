package stack

type element[V any] struct {
	val  V
	next *element[V]
}

// Stack represents a stack data structure. It is a LIFO (last in first out) data
// structure, which means that the last element added to the stack will be the
// first one to be removed.
type Stack[V any] struct {
	head *element[V]
	len  int
}

// New returns a new stack.
func New[V any]() *Stack[V] {
	dummy := &element[V]{}
	return &Stack[V]{
		head: dummy,
		len:  0,
	}
}

// Len returns the number of elements in the stack.
func (s *Stack[V]) Len() int {
	return s.len
}

// Push adds an element to the stack.
func (s *Stack[V]) Push(v V) {
	e := &element[V]{val: v}
	e.next = s.head
	s.head = e
	s.len++
}

// Peek returns the last element added to the stack without removing it.
func (s *Stack[V]) Peek() *V {
	if s.len == 0 {
		return nil
	}
	return &s.head.val
}

// Pop removes and returns the last element added to the stack.
func (s *Stack[V]) Pop() *V {
	if s.len == 0 {
		return nil
	}
	e := s.head
	s.head = e.next
	s.len--
	e.next = nil
	return &e.val
}
