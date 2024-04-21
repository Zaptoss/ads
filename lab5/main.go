package main

import (
	"bufio"
	"fmt"
	"lab5/bst/bytebst"
	"lab5/bst/stringbst"
	"os"
)

func userInputRead() string {
	in := bufio.NewReader(os.Stdin)
	userInput, _, err := in.ReadLine()
	if err != nil {
		panic(fmt.Errorf("error when scanln: %w", err))
	}
	return string(userInput)
}

func deleteAllKeyTask1(tree bytebst.ByteBST, key int) {
	node := tree.Search(key)
	for !node.IsNil() {
		tree.Delete(key)
		node = tree.Search(key)
	}
}

func deleteAllKeyTask2(tree stringbst.StringBST, key int) {
	node := tree.Search(key)
	for !node.IsNil() {
		tree.Delete(key)
		node = tree.Search(key)
	}
}

func rmDuplicateTask1(tree bytebst.ByteBST) {
	x := tree.Min()
	for !x.IsNil() {
		y := x.Successor()
		if y.IsNil() {
			return
		}

		if y.Key() == x.Key() {
			x = x.Predecessor()
			if x.IsNil() {
				x = tree.Min()
			}

			deleteAllKeyTask1(tree, y.Key())
			continue
		}

		x = y
	}
}

func PostorderTask1(tree bytebst.ByteBST) {
	if tree.Root().IsNil() {
		return
	}
	postorderTask1(tree.Root())
	fmt.Print("\n")
}

func postorderTask1(node bytebst.ByteBSTNode) {
	if !node.Left().IsNil() {
		postorderTask1(node.Left())
	}

	if !node.Right().IsNil() {
		postorderTask1(node.Right())
	}

	fmt.Print(node)
}

type NODE struct {
	data  string
	left  *NODE
	right *NODE
}

func postorderTask3(root *NODE) {
	if root.left != nil {
		postorderTask3(root.left)
	}

	if root.right != nil {
		postorderTask3(root.right)
	}

	fmt.Print(root.data)
}

func preorderTask3(root *NODE) {
	fmt.Print(root.data)

	if root.left != nil {
		preorderTask3(root.left)
	}

	if root.right != nil {
		preorderTask3(root.right)
	}
}

func inorderTask3(root *NODE) {
	brackes := root.left != nil || root.right != nil

	if brackes {
		fmt.Print("(")
	}

	if root.left != nil {
		inorderTask3(root.left)
	}

	fmt.Print(root.data)

	if root.right != nil {
		inorderTask3(root.right)
	}

	if brackes {
		fmt.Print(")")
	}
}

func main() {
	// Task 1
	// Побудувати двійкове дерево пошуку з букв рядка,
	// що вводиться. Вивести його на екран у вигляді
	// дерева. Знайти букви, що зустрічаються більше
	// одного разу. Видалити з дерева ці літери.
	// Вивести елементи дерева, що залишилися, при
	// його постфіксному обході.
	charTree := bytebst.NewByteBST()
	input := userInputRead()
	for _, char := range input {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			charTree.Insert(int(char), byte(char))
		}
	}

	rmDuplicateTask1(charTree)
	PostorderTask1(charTree)
	fmt.Println(charTree)

	// Task 2
	// Побудувати двійкове дерево пошуку, в вершинах якого
	// знаходяться слова з текстового файлу. Вивести його на
	// екран у вигляді дерева. Визначити кількість вершин
	// дерева, що містять слова, які починаються на зазначену
	// букву. Видалити з дерева ці вершини.
	stringTree := stringbst.NewStringBST()

	file, err := os.Open("user.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		stringTree.Insert(int(text[0]), text)
	}
	// fmt.Println(stringTree)
	// fmt.Println(stringTree.Count(int('C')))
	deleteAllKeyTask2(stringTree, int('C'))

	// Task 3
	//Побудувати і вивести на екран бінарне дерево наступного виразу:
	// 9 +  8 * (7 + (6 * (5 + 4) – (3 – 2)) +1) .
	// Написати функції постфіксного, інфіксного та префіксного обходу
	// дерева і вивести відповідні вирази на екран.
	res := &NODE{data: "+"}
	res.left = &NODE{data: "9"}
	res.right = &NODE{data: "*"}
	x := res.right
	x.left = &NODE{data: "8"}
	x.right = &NODE{data: "+"}
	x = x.right
	x.left = &NODE{data: "7"}
	x.right = &NODE{data: "+"}
	x = x.right
	x.right = &NODE{data: "1"}
	x.left = &NODE{data: "-"}
	x = x.left
	x.left = &NODE{data: "*"}
	x.left.left = &NODE{data: "6"}
	x.left.right = &NODE{data: "+"}
	x.left.right.left = &NODE{data: "5"}
	x.left.right.right = &NODE{data: "4"}
	x.right = &NODE{data: "-"}
	x.right.left = &NODE{data: "3"}
	x.right.right = &NODE{data: "2"}
	fmt.Print("Postorder: ")
	postorderTask3(res)
	fmt.Println()

	fmt.Print("Preorder: ")
	preorderTask3(res)
	fmt.Println()

	fmt.Print("Inorder: ")
	inorderTask3(res)
	fmt.Println()
}
