package task1

// Піднести до додатного цілого степеня дійсне ненульове число.

func RecursivePow(base float64, power uint) float64 {
	if power < 1 {
		return 0
	} else if power == 1 {
		return base
	} else {
		return base * RecursivePow(base, power-1)
	}
}
