package task8

// Обчислити суму арифметичної прогресії за допомогою рекурсії.

func RecursiveArithmeticProgressionSum(firstElement, step float64, elementNumber uint) float64 {
	if elementNumber == 1 {
		return firstElement
	} else if elementNumber == 0 {
		return 0
	}
	return firstElement + RecursiveArithmeticProgressionSum(firstElement+step, step, elementNumber-1)
}
