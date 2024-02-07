package task12

// Рекурсивно визначити, чи є задане число ступенем двійки.

func RecursiveCheckPowerOfTwo(number int) bool {

	if number <= 0 {
		return false
	}

	return recursiveCheckPowerOfTwo(number)
}

func recursiveCheckPowerOfTwo(number int) bool {
	if number == 1 {
		return true
	} else if number%2 == 1 {
		return false
	}

	return recursiveCheckPowerOfTwo(number / 2)
}
