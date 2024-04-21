package bytebst

import (
	"fmt"
	"lab5/queue"
)

type ByteBST interface {
	Root() ByteBSTNode
	Search(key int) ByteBSTNode
	Max() ByteBSTNode
	Min() ByteBSTNode

	Insert(key int, data byte) ByteBSTNode
	Delete(key int)

	Height() int
	String() string

	raw() *byteBST
}

type ByteBSTNode interface {
	Data() byte
	Key() int

	Parent() ByteBSTNode
	Left() ByteBSTNode
	Right() ByteBSTNode

	Successor() ByteBSTNode
	Predecessor() ByteBSTNode

	String() string
	IsNil() bool

	raw() *byteBSTNode
}

type byteBST struct {
	root *byteBSTNode
}

type byteBSTNode struct {
	key    int
	data   byte
	parent *byteBSTNode
	left   *byteBSTNode
	right  *byteBSTNode
}

func NewByteBST() ByteBST {
	return &byteBST{}
}

func (b *byteBST) Root() ByteBSTNode {
	return b.root
}

func (b *byteBST) Insert(key int, data byte) ByteBSTNode {
	node := &byteBSTNode{key: key, data: data}
	if b.root == nil {
		b.root = node
		return node
	}

	parent, target := search(b.root, key)
	for target != nil && target.right != nil {
		parent, target = search(target.right, key)
	}

	if target != nil {
		node.parent = target
		target.right = node
		return node
	}

	node.parent = parent
	if key < parent.key {
		parent.left = node
		return node
	}

	parent.right = node
	return node
}

func (b *byteBST) Delete(key int) {
	_, node := search(b.root, key)
	if node == nil {
		return
	}

	nodeToDel := min(node.right)
	if nodeToDel != nil {
		node.key = nodeToDel.key
		node.data = nodeToDel.data

		if nodeToDel.right != nil {
			nodeToDel.right.parent = nodeToDel.parent
		}

		if nodeToDel == node.right {
			node.right = nodeToDel.right
			return
		}

		nodeToDel.parent.left = nodeToDel.right
		return
	}

	if node.left != nil {
		node.left.parent = node.parent
	}

	if node.parent != nil {
		if node.key < node.parent.key {
			node.parent.left = node.left
			return
		}

		node.parent.right = node.left
		return
	}

	b.root = node.left
}

func (b *byteBST) Height() (height int) {
	node := b.root
	if node == nil {
		return 0
	}

	queues := []*queue.Queue{queue.NewQueue(), queue.NewQueue()}
	for {
		if node.left != nil {
			queues[1].Push(node.left)
		}

		if node.right != nil {
			queues[1].Push(node.right)
		}

		if queues[0].IsEmpty() {
			queues[0], queues[1] = queues[1], queues[0]
			height += 1
		}

		if queues[0].IsEmpty() {
			height -= 1
			break
		}

		node = queues[0].Pop().(*byteBSTNode)
	}

	return height
}

func (b *byteBST) Search(key int) (node ByteBSTNode) {
	_, node = search(b.root, key)
	return node
}

func search(root *byteBSTNode, key int) (parent, target *byteBSTNode) {
	target = root
	parent = target
	for target != nil {
		if target.key == key {
			break
		}

		parent = target

		if target.key > key {
			target = target.left
			continue
		}

		target = target.right
	}

	return parent, target
}

func (b *byteBST) Max() ByteBSTNode {
	return max(b.root)
}

func max(node *byteBSTNode) *byteBSTNode {
	if node == nil {
		return nil
	}

	for node.right != nil {
		node = node.right
	}

	return node
}

func (b *byteBST) Min() ByteBSTNode {
	return min(b.root)
}

func min(node *byteBSTNode) *byteBSTNode {
	if node == nil {
		return nil
	}

	for node.left != nil {
		node = node.left
	}

	return node
}

type byteNode struct {
	pos  int
	deep int
	node *byteBSTNode
}

func preorderForString(node *byteBSTNode, deep int, pos int, queue *queue.Queue) {
	queue.Push(byteNode{pos: pos, deep: deep, node: node})
	if node.left != nil {
		preorderForString(node.left, deep+1, -1, queue)
	}

	if node.right != nil {
		preorderForString(node.right, deep+1, 1, queue)
	}
}

func (b *byteBST) String() string {
	if b.root == nil {
		return "Tree empty"
	}

	queue := queue.NewQueue()
	preorderForString(b.root, 0, 0, queue)

	var result string
	for !queue.IsEmpty() {
		element := queue.Pop().(byteNode)
		if element.deep == 0 {
			result += fmt.Sprintf("Root(%d): %s\n", element.node.key, fmt.Sprint(element.node))
			continue
		}
		result += "|"
		for i := 0; i < element.deep; i++ {
			result += "----"
		}

		if element.pos == -1 {
			result += fmt.Sprintf("Left(%d): %s\n", element.node.key, fmt.Sprint(element.node))
			continue
		}

		result += fmt.Sprintf("Right(%d): %s\n", element.node.key, fmt.Sprint(element.node))

	}

	return result
}

func (b *byteBST) raw() *byteBST {
	return b
}

// ByteBSTNode implementation

func (n *byteBSTNode) String() string {
	return string(n.data)
}

func (n *byteBSTNode) Key() int {
	return n.key
}

func (n *byteBSTNode) Data() byte {
	return n.data
}

func (n *byteBSTNode) Parent() ByteBSTNode {
	return n.parent
}

func (n *byteBSTNode) Left() ByteBSTNode {
	return n.left
}

func (n *byteBSTNode) Right() ByteBSTNode {
	return n.right
}

func (n *byteBSTNode) Successor() ByteBSTNode {
	return n.successor()
}

func (n *byteBSTNode) successor() *byteBSTNode {
	x := n
	if x.right != nil {
		return min(x.right)
	}

	y := x.parent
	for y != nil && x == y.right {
		x = y
		y = y.parent
	}

	return y
}

func (n *byteBSTNode) Predecessor() ByteBSTNode {
	return n.predecessor()
}

func (n *byteBSTNode) predecessor() *byteBSTNode {
	x := n
	if x.left != nil {
		return max(x.left)
	}

	y := x.parent
	for y != nil && x == y.left {
		x = y
		y = y.parent
	}

	return y
}

func (n *byteBSTNode) IsNil() bool {
	return n == nil
}

func (n *byteBSTNode) raw() *byteBSTNode {
	return n
}
