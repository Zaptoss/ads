package queue

import (
	"fmt"
	"strings"
)

type node struct {
	data any
	next *node
}

type Queue struct {
	head   *node
	tail   *node
	lenght int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(element any) {
	newNode := node{data: element}
	if q.lenght == 0 {
		q.head = &newNode
		q.tail = &newNode
		q.lenght = 1
		return
	}

	q.tail.next = &newNode
	q.tail = &newNode
	q.lenght += 1
}

func (q *Queue) Pop() (element any) {
	if q.lenght == 0 {
		return nil
	}

	element = q.head.data
	q.head = q.head.next
	q.lenght -= 1
	if q.lenght == 0 {
		q.tail = nil
	}

	return
}

func (q *Queue) Length() int {
	return q.lenght
}

func (q *Queue) IsEmpty() bool {
	return q.lenght == 0
}

func (q *Queue) String() string {
	var res []string
	cur := q.head
	for cur != nil {
		res = append(res, fmt.Sprint(cur.data))
		cur = cur.next
	}
	return strings.Join(res, ", ")
}
