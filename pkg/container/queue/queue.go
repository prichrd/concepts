package queue

type element[V any] struct {
	val  V
	next *element[V]
}

// Queue represents a queue data structure. It is a FIFO (first in first out) data
// structure, which means that the first element added to the queue will be the
// first one to be removed.
type Queue[V any] struct {
	head *element[V]
	tail *element[V]
	len  int
}

// New returns a new queue.
func New[V any]() *Queue[V] {
	dummy := &element[V]{}
	return &Queue[V]{
		head: dummy,
		tail: dummy,
		len:  0,
	}
}

// Len returns the number of elements in the queue.
func (q *Queue[V]) Len() int {
	return q.len
}

// Push adds an element to the queue.
func (q *Queue[V]) Push(v V) {
	e := &element[V]{val: v}
	q.tail.next = e
	q.tail = e
	q.len++
}

// Peek returns the first element in the queue.
func (q *Queue[V]) Peek() *V {
	if q.len == 0 {
		return nil
	}
	return &q.head.next.val
}

// Pop removes and returns the first element in the queue.
func (q *Queue[V]) Pop() *V {
	if q.len == 0 {
		return nil
	}
	e := q.head.next
	q.head.next = e.next
	q.len--
	e.next = nil
	return &e.val
}
