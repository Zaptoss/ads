package task10

// Знайти найменше спільне кратне (НСК) двох цілих чисел за допомогою рекурсії.

func RecursiveLCM(a, b int) int {
	if a == 0 || b == 0 {
		return -1
	}
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return recursiveLCM(a, b, a, b)
}

func recursiveLCM(a, b int, c, d int) int {
	if c > d {
		d += b
	} else if d > c {
		c += a
	} else {
		return c
	}

	return recursiveLCM(a, b, c, d)
}
