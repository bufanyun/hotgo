package queue

import (
	"container/list"
	"sync"
)

type Queue struct {
	l *list.List
	m sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{l: list.New()}
}

func (q *Queue) LPush(v interface{}) {
	if v == nil {
		return
	}
	q.m.Lock()
	defer q.m.Unlock()
	q.l.PushFront(v)
}

func (q *Queue) RPush(v interface{}) {
	if v == nil {
		return
	}
	q.m.Lock()
	defer q.m.Unlock()
	q.l.PushBack(v)
}

func (q *Queue) LPop() interface{} {
	q.m.Lock()
	defer q.m.Unlock()

	element := q.l.Front()
	if element == nil {
		return nil
	}

	q.l.Remove(element)
	return element.Value
}

func (q *Queue) RPop() interface{} {
	q.m.Lock()
	defer q.m.Unlock()

	element := q.l.Back()
	if element == nil {
		return nil
	}

	q.l.Remove(element)
	return element.Value
}

func (q *Queue) Len() int {
	return q.l.Len()
}
