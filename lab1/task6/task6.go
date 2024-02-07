package task6

// Рекурсивно обчислити суму цифр заданого натурального числа.

func RecursiveDigitsSum(number uint) uint {
	if number == 0 {
		return 0
	}
	return number%10 + RecursiveDigitsSum(number/10)
}
