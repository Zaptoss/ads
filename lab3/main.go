package main

import (
	"bufio"
	"fmt"
	dd "lab3/dynamicarray"
	ll "lab3/linkedlist"
	"os"
	"strconv"
	"strings"
)

func validateString(s string) (stringValidation bool) {
	stringValidation = true
	for _, letter := range s {
		if letter >= 'a' && letter <= 'z' {
			continue
		}

		if letter >= 'A' && letter <= 'Z' {
			continue
		}

		if letter == ' ' {
			continue
		}
		stringValidation = false
	}
	return
}

func userInputRead() string {
	in := bufio.NewReader(os.Stdin)
	userInput, _, err := in.ReadLine()
	if err != nil {
		panic(fmt.Errorf("error when scanln: %w", err))
	}
	return string(userInput)
}

func testLinkedList() {
	fmt.Println(`Test linked list
0 - Exit
1 - Add nodes(You can write only letters and whitespaces which would divide string by nodes) (! You can use this option only once for initialize LL !)
2 - Remove (remove node by index)
3 - Length
4 - Data (Data specified node index)
5 - List (Show list)
6 - Add after(you enter index and string and it string insert after specified index)
7 - Add before(you enter index and string and it string insert before specified index)`)

	linkedList := ll.NewDoublyLinkedList()

	for {
		var userInput string

		fmt.Print("Enter option: ")
		userInput = userInputRead()

		opt, err := strconv.ParseUint(userInput, 10, 64)
		if err != nil {
			fmt.Printf("Error when parse input: %s\n", err)
			continue
		}

		switch opt {
		case 0:
			fmt.Println("Goodbay!")
			return

		case 1:
			if linkedList.Length() > 0 {
				fmt.Println("Linked List already initialized")
				break
			}

			fmt.Print("Enter string: ")
			userInput = userInputRead()
			if !validateString(userInput) {
				fmt.Println("Incorrect format")
				break
			}

			for _, item := range strings.Split(userInput, " ") {
				linkedList.InsertEnd(ll.NewDoublyLinkedNode(item))
			}

		case 2:
			fmt.Print("Enter index: ")
			userInput = userInputRead()
			node, err := strconv.ParseUint(userInput, 10, 64)
			if err != nil {
				fmt.Printf("Error when parse input: %s\n", err)
				break
			}

			if node >= uint64(linkedList.Length()) {
				fmt.Println("out of list")
				break
			}

			rmNode := linkedList.First()
			for i := 0; i < int(node); i++ {
				rmNode = rmNode.Next()
			}

			linkedList.Remove(rmNode)

		case 3:
			fmt.Printf("Length of list: %d\n", linkedList.Length())

		case 4:
			fmt.Print("Enter index: ")
			userInput = userInputRead()
			node, err := strconv.ParseUint(userInput, 10, 64)
			if err != nil {
				fmt.Printf("Error when parse input: %s\n", err)
				break
			}

			if node >= uint64(linkedList.Length()) {
				fmt.Println("out of list")
				break
			}

			outNode := linkedList.First()
			for i := 0; i < int(node); i++ {
				outNode = outNode.Next()
			}

			fmt.Printf("Data %d node: %s", node, outNode.Data())

		case 5:
			fmt.Println(linkedList)

		case 6:
			fmt.Print("Enter index: ")
			userInput = userInputRead()
			node, err := strconv.ParseUint(userInput, 10, 64)
			if err != nil {
				fmt.Printf("Error when parse input: %s\n", err)
				break
			}

			if node >= uint64(linkedList.Length()) {
				fmt.Println("out of list")
				break
			}

			aftNode := linkedList.First()
			for i := 0; i < int(node); i++ {
				aftNode = aftNode.Next()
			}

			fmt.Print("Enter string: ")
			userInput = userInputRead()
			if !validateString(userInput) {
				fmt.Println("Incorrect format")
				break
			}

			for _, item := range strings.Split(userInput, " ") {
				linkedList.InsertAfter(aftNode, ll.NewDoublyLinkedNode(item))
				aftNode = aftNode.Next()
			}

		case 7:
			fmt.Print("Enter index: ")
			userInput = userInputRead()
			node, err := strconv.ParseUint(userInput, 10, 64)
			if err != nil {
				fmt.Printf("Error when parse input: %s\n", err)
				break
			}

			if node >= uint64(linkedList.Length()) {
				fmt.Println("out of list")
				break
			}

			bfNode := linkedList.First()
			for i := 0; i < int(node); i++ {
				bfNode = bfNode.Next()
			}

			fmt.Print("Enter string: ")
			userInput = userInputRead()
			if !validateString(userInput) {
				fmt.Println("Incorrect format")
				break
			}

			for _, item := range strings.Split(userInput, " ") {
				linkedList.InsertBefore(bfNode, ll.NewDoublyLinkedNode(item))
				bfNode = bfNode.Next()
			}

		case 8:
			// Traversal
			startNode := linkedList.First()
			for startNode != nil {
				data := startNode.Data()
				if len(data)%2 == 1 {
					if len(data) == 1 {
						fmt.Printf(",")
					} else {
						fmt.Printf("%s,", data[1:len(data)-1])
					}
				}
				startNode = startNode.Next()
			}
			fmt.Println()

		default:
			fmt.Println("Incorrect input")
		}

	}
}

func testDynamicArray() {
	fmt.Println(`Test dynamic array
0 - Exit
1 - Append
2 - Pop
3 - Length
4 - Get
5 - Set
6 - Insert after
7 - Insert before
8 - List
9 - Remove`)

	dynamicArray := dd.NewDynamicArray()

	for {
		var userInput string

		fmt.Print("Enter option: ")
		userInput = userInputRead()

		opt, err := strconv.ParseUint(userInput, 10, 64)
		if err != nil {
			fmt.Printf("Error when parse input: %s\n", err)
			continue
		}

		switch opt {
		case 0:
			fmt.Println("Goodbay!")
			return

		case 1:
			fmt.Print("Enter string: ")
			userInput = userInputRead()
			if !validateString(userInput) {
				fmt.Println("Incorrect format")
				break
			}

			for _, item := range strings.Split(userInput, " ") {
				dynamicArray.Append(item)
			}

		case 2:
			data, err := dynamicArray.Pop()
			if err != nil {
				fmt.Printf("Error when pop element: %s\n", err)
				break
			}

			fmt.Printf("Data: %s\n", data)

		case 3:
			fmt.Printf("Length of array: %d\n", dynamicArray.Length())

		case 4:
			fmt.Print("Enter index: ")
			userInput = userInputRead()
			index, err := strconv.ParseUint(userInput, 10, 64)
			if err != nil {
				fmt.Printf("Error when parse input: %s\n", err)
				break
			}

			data, err := dynamicArray.Get(uint(index))
			if err != nil {
				fmt.Printf("Error when get element: %s\n", err)
				break
			}

			fmt.Printf("Data: %s\n", data)

		case 5:
			fmt.Print("Enter index: ")
			userInput = userInputRead()
			index, err := strconv.ParseUint(userInput, 10, 64)
			if err != nil {
				fmt.Printf("Error when parse input: %s\n", err)
				break
			}

			fmt.Print("Enter string: ")
			userInput = userInputRead()
			if !validateString(userInput) {
				fmt.Println("Incorrect format")
				break
			}

			if len(strings.Split(userInput, " ")) == 1 {
				err = dynamicArray.Set(uint(index), userInput)
				if err != nil {
					fmt.Printf("Error when set value: %s\n", err)
				}
				break
			}

			fmt.Println("Input must be one word")

		case 6:
			fmt.Print("Enter index: ")
			userInput = userInputRead()
			index, err := strconv.ParseUint(userInput, 10, 64)
			if err != nil {
				fmt.Printf("Error when parse input: %s\n", err)
				break
			}

			fmt.Print("Enter string: ")
			userInput = userInputRead()
			if !validateString(userInput) {
				fmt.Println("Incorrect format")
				break
			}

			for _, item := range strings.Split(userInput, " ") {
				err = dynamicArray.InsertAfter(uint(index), item)
				if err != nil {
					fmt.Printf("Error when insert string: %s\n", err)
					break
				}
				index += 1
			}
		case 7:
			fmt.Print("Enter index: ")
			userInput = userInputRead()
			index, err := strconv.ParseUint(userInput, 10, 64)
			if err != nil {
				fmt.Printf("Error when parse input: %s\n", err)
				break
			}

			fmt.Print("Enter string: ")
			userInput = userInputRead()
			if !validateString(userInput) {
				fmt.Println("Incorrect format")
				break
			}

			for _, item := range strings.Split(userInput, " ") {
				err = dynamicArray.InsertBefore(uint(index), item)
				if err != nil {
					fmt.Printf("Error when insert string: %s\n", err)
					break
				}
			}

		case 8:
			fmt.Println(dynamicArray)

		case 9:
			fmt.Print("Enter index: ")
			userInput = userInputRead()
			index, err := strconv.ParseUint(userInput, 10, 64)
			if err != nil {
				fmt.Printf("Error when parse input: %s\n", err)
				break
			}

			err = dynamicArray.Remove(uint(index))
			if err != nil {
				fmt.Printf("Error when get element: %s\n", err)
			}

		case 10:
			for i := uint(0); i < dynamicArray.Length(); i++ {
				if s, _ := dynamicArray.Get(i); len(s)%2 == 1 {
					fmt.Printf("%s,\n", s[1:len(s)-1])
				}
			}

		default:
			fmt.Println("Incorrect input")
		}

	}
}

func main() {
	testLinkedList()
	testDynamicArray()
}
