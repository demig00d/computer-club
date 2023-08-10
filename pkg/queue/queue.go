package queue

import (
	"errors"
)

var (
	ErrElementIsAlreadyInQueue = errors.New("element is already in queue")
	ErrEmptyQueue              = errors.New("queue is empty")
)

type Queue[T comparable] struct {
	ts []T
}

func NewQueue[T comparable]() Queue[T] {
	return Queue[T]{
		ts: make([]T, 0),
	}
}

func (q *Queue[T]) Push(t T) error {
	for _, el := range q.ts {
		if el == t {
			return ErrElementIsAlreadyInQueue
		}
	}
	q.ts = append(q.ts, t)
	return nil
}

func (q *Queue[T]) Poll() (T, error) {
	var defaultValue T
	if len(q.ts) < 1 {
		return defaultValue, ErrEmptyQueue
	}

	front := q.ts[0]
	q.ts = q.ts[1:]
	return front, nil
}

func (q Queue[T]) Length() int {
	return len(q.ts)
}

func (q Queue[T]) IsIn(t T) bool {
	for _, tt := range q.ts {
		if tt == t {
			return true
		}

	}
	return false
}

func (q *Queue[T]) Remove(t T) {
	if len(q.ts) < 1 {
		return
	}
	for i, el := range q.ts {
		if el == t {
			copy(q.ts[i:], q.ts[i+1:])
			q.ts = q.ts[:len(q.ts)-1]
			return
		}
	}
}
