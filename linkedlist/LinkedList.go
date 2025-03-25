package linkedlist

type E[T any] struct {
	prv, next *E[T]
	V         T
	l         *LinkedList[T]
}

func (e *E[T]) Next() *E[T] {
	if n := e.next; n != &e.l.root {
		return n
	}
	return nil
}

func (e *E[T]) Prev() *E[T] {
	if p := e.prv; p != &e.l.root {
		return p
	}
	return nil
}

type LinkedList[T any] struct {
	root E[T]
	len  int
}

func New[T any]() *LinkedList[T] {
	l := &LinkedList[T]{}
	l.root.prv = &l.root
	l.root.next = &l.root
	return l
}

func (l *LinkedList[T]) Len() int {
	return l.len
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.len == 0
}

func (l *LinkedList[T]) Clear() {}

func (l *LinkedList[T]) Front() *E[T] {
	if l.IsEmpty() {
		return nil
	}
	return l.root.next
}

func (l *LinkedList[T]) Back() *E[T] {
	if l.IsEmpty() {
		return nil
	}
	return l.root.prv
}

func (l *LinkedList[T]) PushBack(v T) *E[T] {
	e := E[T]{V: v, prv: nil, next: nil}
	return l.insert(&e, l.root.prv)
}

func (l *LinkedList[T]) PushFront(v T) *E[T] {
	e := E[T]{V: v, prv: nil, next: nil}
	return l.insert(&e, &l.root)
}

func (l *LinkedList[T]) Remove(e *E[T]) {
	if l.IsEmpty() {
		return
	}

	e.prv.next = e.next
	e.next.prv = e.prv

	e.prv = nil
	e.next = nil
	e = nil

	l.len--
}

func (l *LinkedList[T]) insert(e *E[T], at *E[T]) *E[T] {
	e.prv = at
	e.next = at.next
	at.next = e
	e.next.prv = e
	l.len++
	e.l = l
	return e
}
