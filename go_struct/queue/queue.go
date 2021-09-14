package queue

import (
	"container/list"
	"fmt"
)

type Queue struct {
	root *list.List
}

func New() (q *Queue) {
	q = &Queue{}
	q.root = list.New()

	return q
}

func (q *Queue) Enqueue(v interface{}) {
	q.root.PushFront(v)
}

func (q *Queue) Dequeue() interface{} {
	return q.root.Remove(q.root.Back())
}

func (q *Queue) Print() {
	for e := q.root.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
