package main

import (
	"fmt"
	linkedqeueue "lab4/LinkedQeueue"
)

func main() {
	queue := linkedqeueue.NewQueue()

	fmt.Println(queue.Length())
	queue.Pop()
	fmt.Println(queue.Length())

	queue.Push(23)
	queue.Push(22)
	queue.Push(151)
	queue.Push(-5)
	queue.Push(13)
	queue.Push(10)
	queue.Push(23)
	queue.Push(22)
	queue.Push(151)
	queue.Push(-5)
	queue.Push(13)
	queue.Push(10)
	fmt.Println(queue.Length())
	fmt.Println(queue)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue)

}
