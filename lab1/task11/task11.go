package task11

// Знайти суму парних або непарних чисел в масиві за допомогою рекурсії.

func RecursiveSumOddArrayNumbers(array []int) int {
	if len(array) == 0 {
		return 0
	}

	if array[0]%2 == 0 {
		return RecursiveSumOddArrayNumbers(array[1:])
	}
	return array[0] + RecursiveSumOddArrayNumbers(array[1:])
}

func RecursiveSumEvenArrayNumbers(array []int) int {
	if len(array) == 0 {
		return 0
	}

	if array[0]%2 == 1 {
		return RecursiveSumEvenArrayNumbers(array[1:])
	}
	return array[0] + RecursiveSumEvenArrayNumbers(array[1:])
}
