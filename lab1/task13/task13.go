package task13

// Знайти суму перших n непарних (або парних) чисел за допомогою рекурсії.

func RecursiveSumFirstEvenNumbers(n uint) uint {
	if n == 0 {
		return 0
	}
	return 2*n + RecursiveSumFirstEvenNumbers(n-1)
}

func RecursiveSumFirstOddNumbers(n uint) uint {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return 2*n - 1 + RecursiveSumFirstOddNumbers(n-1)
}
