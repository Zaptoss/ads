package linkedlist

import (
	"fmt"
	"strings"
)

type DoublyLinkedNode struct {
	next *DoublyLinkedNode
	prev *DoublyLinkedNode
	data string
}

type DoublyLinkedList struct {
	first  *DoublyLinkedNode
	last   *DoublyLinkedNode
	length uint
}

func NewDoublyLinkedNode(data string) *DoublyLinkedNode {
	return &DoublyLinkedNode{data: data}
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (n *DoublyLinkedNode) Data() string {
	return n.data
}

func (n *DoublyLinkedNode) Next() *DoublyLinkedNode {
	return n.next
}

func (n *DoublyLinkedNode) Prev() *DoublyLinkedNode {
	return n.prev
}

func (l *DoublyLinkedList) Length() uint {
	return l.length
}

func (l *DoublyLinkedList) First() *DoublyLinkedNode {
	return l.first
}

func (l *DoublyLinkedList) Last() *DoublyLinkedNode {
	return l.last
}

func (l *DoublyLinkedList) InsertAfter(node, newNode *DoublyLinkedNode) {
	l.length += 1

	newNode.prev = node
	newNode.next = node.next
	node.next = newNode

	if newNode.next == nil {
		l.last = newNode
		return
	}

	newNode.next.prev = newNode
}

func (l *DoublyLinkedList) InsertBefore(node, newNode *DoublyLinkedNode) {
	l.length += 1

	newNode.next = node
	newNode.prev = node.prev
	node.prev = newNode

	if newNode.prev == nil {
		l.first = newNode
		return
	}

	newNode.prev.next = newNode
}

func (l *DoublyLinkedList) InsertBeginning(node *DoublyLinkedNode) {
	if l.first == nil {
		l.length += 1
		l.first = node
		l.last = node
		node.prev = nil
		node.next = nil
		return
	}

	l.InsertBefore(l.first, node)
}

func (l *DoublyLinkedList) InsertEnd(node *DoublyLinkedNode) {
	if l.last == nil {
		l.InsertBeginning(node)
		return
	}

	l.InsertAfter(l.last, node)
}

func (l *DoublyLinkedList) Remove(node *DoublyLinkedNode) (err error) {

	if l.length == 0 {
		return fmt.Errorf("nothing to remove from Doubluy Linked List")
	}

	if l.length == 1 { // node are single in linked list; it mean that prev and next is nil
		l.first = nil // in empty LL first - nil
		l.last = nil  // in empty LL last - nil
		l.length = 0
		return
	}

	l.length -= 1

	if node.prev == nil { // prev not exists and next exists
		l.first = node.next // next node become first
		l.first.prev = nil  // prev node for first is nil
		return
	}

	if node.next == nil { // prev exists and next not exists
		l.last = node.prev // prev node become last
		l.last.next = nil  // next node for last is nil
		return
	}

	node.prev.next = node.next // next for prev node point to next for current
	node.next.prev = node.prev // prev for next node point to prev for current
	return
}

func (l *DoublyLinkedList) String() string {
	if l.length == 0 {
		return "List is empty"
	}

	var dataList []string
	node := l.first

	for node != nil {
		dataList = append(dataList, fmt.Sprintf("(%s)", node.data))
		node = node.next
	}
	return "Head -> " + strings.Join(dataList, " <-> ") + " <- Tail"
}
