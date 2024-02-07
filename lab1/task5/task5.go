package task5

// Визначити, чи є задане слово паліндромом (слово, яке читається однаково зліва направо і справа наліво).

func RecursivePalindromCheck(word string) bool {
	wordLenght := len(word)
	if wordLenght < 2 {
		return true
	}

	if word[0] == word[wordLenght-1] {
		return RecursivePalindromCheck(word[1 : wordLenght-1])
	}

	return false
}
