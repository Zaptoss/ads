package task7

import "math"

// Визначити, чи є задане число простим за допомогою рекурсії.

func RecursivePrimeTest(number int) bool {
	if number < 0 {
		number = -number
	}

	maxNumber := int(math.Sqrt(float64(number)))

	return recursivePrimeTest(number, 2, maxNumber)
}

func recursivePrimeTest(number, testNumber, maxNumber int) bool {
	if number == 0 || number == 1 {
		return false
	}
	if number == 2 || number == 3 {
		return true
	}

	if testNumber <= maxNumber {
		if number%testNumber == 0 {
			return false
		}
		return recursivePrimeTest(number, testNumber+1, maxNumber)
	}
	return true
}
