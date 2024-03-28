package arrayqueue

import (
	"fmt"
	"strings"
)

// Lenght can't be greather than size-1
type Queue struct {
	array  []int
	head   int
	tail   int
	size   int
	length int
}

func NewQueue(size int) (*Queue, error) {
	if size <= 0 {
		return nil, fmt.Errorf("queue size must be greater than 0")
	}

	return &Queue{array: make([]int, size), size: size}, nil
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Length() int {
	return q.length
}

func (q *Queue) Push(element int) {
	if q.length == q.size-1 {
		return
	}

	q.array[q.tail] = element
	q.tail = (q.tail + 1) % q.size
	q.length += 1
}

func (q *Queue) Pop() (element int) {
	if q.length == 0 {
		return -1
	}

	element = q.array[q.head]
	q.head = (q.head + 1) % q.size
	q.length -= 1
	return
}

func (q *Queue) IsEmpty() bool {
	return q.length == 0
}

func (q *Queue) String() string {
	if q.length == 0 {
		return "Queue empty"
	}

	result := make([]string, q.length)
	cur := q.head
	i := 0
	for cur != q.tail {
		result[i] = fmt.Sprint(q.array[cur])
		cur = (cur + 1) % q.size
		i += 1
	}
	return strings.Join(result, ", ")
}
