package task4

// Рекурсивно знайти суму перших n натуральних чисел.

func RecursiveNaturalSum(number uint) uint {
	if number <= 1 {
		return number
	}
	return number + RecursiveNaturalSum(number-1)
}
