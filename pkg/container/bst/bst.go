package bst

import (
	"golang.org/x/exp/constraints"
)

// Tree represents a binary search tree (BST) structure. This structure
// facilitates the traversal of elements in order.
type Tree[K constraints.Ordered, V any] struct {
	root *tree[K, V]
}

// New returns a new Tree.
func New[K constraints.Ordered, V any]() *Tree[K, V] {
	return &Tree[K, V]{}
}

// Insert adds an element to the BST.
func (t *Tree[K, V]) Insert(key K, val V) {
	tree := &tree[K, V]{
		key: key,
		val: val,
	}
	if t.root == nil {
		t.root = tree
		return
	}
	t.root.insert(tree)
}

// Remove deletes an element from the BST. The performance of the
// delete operation highly depends on the order of insertion, as a
// BST isn't a balanced structure. In cases where the element is not
// present in the tree, a zero value with a `false` boolean will be
// returned. If the object is found, it is removed from the structure
// and returned.
func (t *Tree[K, V]) Remove(k K) (V, bool) {
	if t == nil {
		var zero V
		return zero, false
	}
	f := t.root.remove(k, nil)
	if f == nil {
		var zero V
		return zero, false
	}
	return f.val, true
}

// Find finds an element from the BST. The performance of the find
// operation highly depends on the order of insertion, as a BST isn't
// a balanced structure. In cases where the element is not present in
// the tree, a zero value with a `false` boolean will be returned.
func (t *Tree[K, V]) Find(k K) (V, bool) {
	f := t.root.find(k)
	if f == nil {
		var zero V
		return zero, false
	}
	return f.val, true
}
func (t *Tree[K, V]) Walk(fn func(K, V)) {
	t.root.walk(fn)
}

type tree[K constraints.Ordered, V any] struct {
	key   K
	val   V
	left  *tree[K, V]
	right *tree[K, V]
}

func (t *tree[K, V]) insert(nt *tree[K, V]) {
	switch {
	case nt.key == t.key:
		// This implementation doesn't accept duplicates. The
		// value is updated if the key is found during insert.
		t.val = nt.val
	case nt.key < t.key && t.left == nil:
		t.left = nt
	case nt.key < t.key:
		t.left.insert(nt)
	case nt.key > t.key && t.right == nil:
		t.right = nt
	default:
		t.right.insert(nt)
	}
}

func (t *tree[K, V]) find(key K) *tree[K, V] {
	switch {
	case t == nil:
		return nil
	case key == t.key:
		return t
	case key < t.key:
		return t.left.find(key)
	case key > t.key:
		return t.right.find(key)
	default:
		return nil
	}
}

func (t *tree[K, V]) swap(with, parent *tree[K, V]) *tree[K, V] {
	if parent == nil {
		return nil
	}
	if t == parent.left {
		tmp := parent.left
		parent.left = with
		return tmp
	}
	tmp := parent.right
	parent.right = with
	return tmp
}

func (t *tree[K, V]) max(parent *tree[K, V]) (*tree[K, V], *tree[K, V]) {
	if t.right == nil {
		return parent, t
	}
	return t.right.max(t)
}

func (t *tree[K, V]) remove(key K, parent *tree[K, V]) *tree[K, V] {
	switch {
	case t == nil:
		return nil
	case key < t.key:
		return t.left.remove(key, t)
	case key > t.key:
		return t.right.remove(key, t)
	case t.left == nil && t.right == nil: // The matching node is a leaf
		return t.swap(nil, parent)
	case t.left != nil && t.right == nil: // Matching node has one child (left)
		return t.swap(t.left, parent)
	case t.right != nil && t.left == nil: // Matching node has one child (right)
		return t.swap(t.right, parent)
	default: // Matching node has 2 children (left and right)
		repParent, rep := t.left.max(t)
		t.key, t.val, rep.key, rep.val = rep.key, rep.val, t.key, t.val
		return rep.swap(nil, repParent)
	}
}

func (t *tree[K, V]) walk(fn func(K, V)) {
	if t == nil {
		return
	}
	t.left.walk(fn)
	fn(t.key, t.val)
	t.right.walk(fn)
}
