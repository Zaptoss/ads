package task9

// Визначити, чи є задана послідовність чисел арифметичною прогресією за допомогою рекурсії.

func RecursiveCheckIfArithmeticProgression(elements ...float64) bool {
	if len(elements) < 3 {
		return true
	}
	if elements[0]-elements[1] != elements[1]-elements[2] {
		return false
	}

	return RecursiveCheckIfArithmeticProgression(elements[1:]...)
}
