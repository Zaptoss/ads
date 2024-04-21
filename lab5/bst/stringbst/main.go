package stringbst

import (
	"fmt"
	"lab5/queue"
)

type StringBST interface {
	Root() StringBSTNode
	Search(key int) StringBSTNode
	Max() StringBSTNode
	Min() StringBSTNode

	Insert(key int, data string) StringBSTNode
	Delete(key int)

	Height() int
	String() string
	Count(key int) int

	raw() *stringBST
}

type StringBSTNode interface {
	Data() string
	Key() int

	Parent() StringBSTNode
	Left() StringBSTNode
	Right() StringBSTNode

	Successor() StringBSTNode
	Predecessor() StringBSTNode

	String() string
	IsNil() bool

	raw() *stringBSTNode
}

type stringBST struct {
	root *stringBSTNode
}

type stringBSTNode struct {
	key    int
	data   string
	parent *stringBSTNode
	left   *stringBSTNode
	right  *stringBSTNode
}

func NewStringBST() StringBST {
	return &stringBST{}
}

func (b *stringBST) Root() StringBSTNode {
	return b.root
}

func (b *stringBST) Insert(key int, data string) StringBSTNode {
	node := &stringBSTNode{key: key, data: data}
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

func (b *stringBST) Delete(key int) {
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

func (b *stringBST) Height() (height int) {
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

		node = queues[0].Pop().(*stringBSTNode)
	}

	return height
}

func (b *stringBST) Count(key int) (count int) {
	node := b.root
	if node == nil {
		return 0
	}

	queues := []*queue.Queue{queue.NewQueue(), queue.NewQueue()}
	for {
		if node.key == key {
			count += 1
		}

		if node.left != nil {
			queues[1].Push(node.left)
		}

		if node.right != nil {
			queues[1].Push(node.right)
		}

		if queues[0].IsEmpty() {
			queues[0], queues[1] = queues[1], queues[0]
		}

		if queues[0].IsEmpty() {
			break
		}

		node = queues[0].Pop().(*stringBSTNode)
	}

	return count
}

func (b *stringBST) Search(key int) (node StringBSTNode) {
	_, node = search(b.root, key)
	return node
}

func search(root *stringBSTNode, key int) (parent, target *stringBSTNode) {
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

func (b *stringBST) Max() StringBSTNode {
	return max(b.root)
}

func max(node *stringBSTNode) *stringBSTNode {
	if node == nil {
		return nil
	}

	for node.right != nil {
		node = node.right
	}

	return node
}

func (b *stringBST) Min() StringBSTNode {
	return min(b.root)
}

func min(node *stringBSTNode) *stringBSTNode {
	if node == nil {
		return nil
	}

	for node.left != nil {
		node = node.left
	}

	return node
}

type stringNode struct {
	pos  int
	deep int
	node *stringBSTNode
}

func preorderForString(node *stringBSTNode, deep int, pos int, queue *queue.Queue) {
	queue.Push(stringNode{pos: pos, deep: deep, node: node})
	if node.left != nil {
		preorderForString(node.left, deep+1, -1, queue)
	}

	if node.right != nil {
		preorderForString(node.right, deep+1, 1, queue)
	}
}

func (b *stringBST) String() string {
	if b.root == nil {
		return "Tree empty"
	}

	queue := queue.NewQueue()
	preorderForString(b.root, 0, 0, queue)

	var result string
	for !queue.IsEmpty() {
		element := queue.Pop().(stringNode)
		if element.deep == 0 {
			result += fmt.Sprintf("Root(%d): %s\n", element.node.key, fmt.Sprint(element.node))
			continue
		}
		result += "|"
		for i := 0; i < element.deep; i++ {
			result += "--"
		}

		if element.pos == -1 {
			result += fmt.Sprintf("Left(%d): %s\n", element.node.key, fmt.Sprint(element.node))
			continue
		}

		result += fmt.Sprintf("Right(%d): %s\n", element.node.key, fmt.Sprint(element.node))

	}

	return result
}

func (b *stringBST) raw() *stringBST {
	return b
}

// StringBSTNode implementation

func (n *stringBSTNode) String() string {
	return n.data
}

func (n *stringBSTNode) Key() int {
	return n.key
}

func (n *stringBSTNode) Data() string {
	return n.data
}

func (n *stringBSTNode) Parent() StringBSTNode {
	return n.parent
}

func (n *stringBSTNode) Left() StringBSTNode {
	return n.left
}

func (n *stringBSTNode) Right() StringBSTNode {
	return n.right
}

func (n *stringBSTNode) Successor() StringBSTNode {
	return n.successor()
}

func (n *stringBSTNode) successor() *stringBSTNode {
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

func (n *stringBSTNode) Predecessor() StringBSTNode {
	return n.predecessor()
}

func (n *stringBSTNode) predecessor() *stringBSTNode {
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

func (n *stringBSTNode) IsNil() bool {
	return n == nil
}

func (n *stringBSTNode) raw() *stringBSTNode {
	return n
}
