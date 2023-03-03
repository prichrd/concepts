package list

// Element represents a node in a doubly linked list. It holds
// a value and references to the previous and next nodes.
type Element[V any] struct {
	prev *Element[V]
	next *Element[V]

	val V
}

// Next returns the pointer to the next element.
func (e *Element[V]) Next() *Element[V] {
	return e.next
}

// Prev returns the pointer to the previous element.
func (e *Element[V]) Prev() *Element[V] {
	return e.prev
}

// Val returns the element's value.
func (e *Element[V]) Val() V {
	return e.val
}

// List represents a linked list with forward and backward
// references.
type List[V any] struct {
	front *Element[V]
	back  *Element[V]
	len   int
}

// NewList returns a configured instance of List.
func NewList[V any]() *List[V] {
	return &List[V]{
		len: 0,
	}
}

// Len returns the length of the linked list.
func (l *List[V]) Len() int {
	return l.len
}

// Front returns the first element of the linked list.
func (l *List[V]) Front() *Element[V] {
	return l.front
}

// Back returns the last element of the linked list
func (l *List[V]) Back() *Element[V] {
	return l.back
}

// Remove removes an element from the linked list.
func (l *List[V]) Remove(e *Element[V]) {
	if e.prev != nil {
		e.prev.next = e.next
	} else {
		e.next.prev = nil
		l.front = e.next
	}
	if e.next != nil {
		e.next.prev = e.prev
	} else {
		e.prev.next = nil
		l.back = e.prev
	}
	e.next = nil // prevent memory leaks
	e.prev = nil // prevent memory leaks
	l.len--
}

// MoveToFront moves an element to the front of the list.
func (l *List[V]) MoveToFront(e *Element[V]) {
	if e.prev == nil {
		// Already in front
		return
	}

	e.prev.next = e.next
	if e.next != nil {
		e.next.prev = e.prev
	} else {
		e.prev.next = nil
		l.back = e.prev
	}
	l.front.prev = e

	e.prev = nil
	e.next = l.front
	l.front = e
}

// PushFront inserts an element in front of the list.
func (l *List[V]) PushFront(v V) *Element[V] {
	n := &Element[V]{
		val: v,
	}

	if l.len == 0 {
		l.front = n
		l.back = n
		l.len++
		return n
	}

	l.front.prev = n
	n.next = l.front
	l.front = n
	l.len++
	return n
}

// PushFront inserts an element at the end of the list.
func (l *List[V]) PushBack(v V) *Element[V] {
	n := &Element[V]{
		val: v,
	}

	if l.len == 0 {
		l.front = n
		l.back = n
		l.len++
		return n
	}

	l.back.next = n
	n.prev = l.back
	l.back = n
	l.len++
	return n
}
