package task2

// Знайти НСД двох цілих чисел за алгоритмом Евкліда.

func RecursiveGCD(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	return recursiveGCD(a, b)
}

func recursiveGCD(a, b int) int {
	r := a % b
	if r == 0 {
		return b
	}
	return recursiveGCD(b, r)
}
