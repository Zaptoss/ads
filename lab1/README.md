# LAB1 ADS
#### В даній лабораторній роботі було використано мову програмування GoLang
---
### Task 1
Реалізовано рекурсивну функцію для піднесення дійсного числа до додатнього цілого степеня: 
+ RecursivePow(base float64, power uint) -> float64
```go
func RecursivePow(base float64, power uint) float64 {
	if power < 1 {
		return 0
	} else if power == 1 {
		return base
	} else {
		return base * RecursivePow(base, power-1)
	}
}
```
---
### Task 2
Реалізовано рекурсивну функцію знаходження НСД за алгоритмом Евкліда
+ RecursiveGCD(a, b int) -> uint
```go
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
```
---
### Task 3
Реалізовано рекурсивну функцію для знаходження n-го числа Фібоначі
+ RecursiveFibonachi(number uint) -> uint
```go
func RecursiveFibonachi(number uint) uint {
	if number <= 1 {
		return 1
	}
	return RecursiveFibonachi(number-1) + RecursiveFibonachi(number-2)
}
```
---
### Task 4
Реалізовано рекурсивну функцію для знаходження суми перших n натуральних чисел
+ RecursiveNaturalSum(number uint) -> uint
```go
func RecursiveNaturalSum(number uint) uint {
	if number <= 1 {
		return number
	}
	return number + RecursiveNaturalSum(number-1)
}
```
---
### Task 5
Реалізовано рекурсивну функцію для перевірки слова на паліндром
+ RecursivePalindromCheck(word string) -> bool
```go
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
```
---
### Task 6
Реалізовано рекурсивну функцію для знаходження суми цифр натурального числа
+ RecursiveDigitsSum(number uint) -> uint
```go
func RecursiveDigitsSum(number uint) uint {
	if number == 0 {
		return 0
	}
	return number%10 + RecursiveDigitsSum(number/10)
}
```
---
### Task 7
Реалізовано рекурсивну функцію для перевірки числа на простоту
+ RecursivePrimeTest(number int) -> bool
```go
import "math"

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
```
---
### Task 8
Реалізовано рекурсивну функцію для обчислення суми арифметичної прогресії
+ RecursiveArithmeticProgressionSum(firstElement, step float64, elementNumber uint) -> float64
```go
func RecursiveArithmeticProgressionSum(firstElement, step float64, elementNumber uint) float64 {
	if elementNumber == 1 {
		return firstElement
	} else if elementNumber == 0 {
		return 0
	}
	return firstElement + RecursiveArithmeticProgressionSum(firstElement+step, step, elementNumber-1)
}
```
---
### Task 9
Реалізовано рекурсивну функцію для перевірки послідовності на арифметичну прогресію
+ RecursiveCheckIfArithmeticProgression(elements ...float64) -> bool
```go
func RecursiveCheckIfArithmeticProgression(elements ...float64) bool {
	if len(elements) < 3 {
		return true
	}
	if elements[0]-elements[1] != elements[1]-elements[2] {
		return false
	}

	return RecursiveCheckIfArithmeticProgression(elements[1:]...)
}
```
---
### Task 10
Реалізовано рекурсивну функцію знаходження НСК
+ RecursiveLCM(a, b int) -> int
```go
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
```
---
### Task 11
Реалізовано рекурсивні функції знаходження суми парних та непарних чисел в масиві
+ RecursiveSumEvenArrayNumbers(array []int) -> int
```go
func RecursiveSumEvenArrayNumbers(array []int) int {
	if len(array) == 0 {
		return 0
	}

	if array[0]%2 == 1 {
		return RecursiveSumEvenArrayNumbers(array[1:])
	}
	return array[0] + RecursiveSumEvenArrayNumbers(array[1:])
}
```
+ RecursiveSumOddArrayNumbers(array []int) -> int
```go
func RecursiveSumOddArrayNumbers(array []int) int {
	if len(array) == 0 {
		return 0
	}

	if array[0]%2 == 0 {
		return RecursiveSumOddArrayNumbers(array[1:])
	}
	return array[0] + RecursiveSumOddArrayNumbers(array[1:])
}
```
---
### Task 12
Реалізовано рекурсивну функцію для перевірки, чи є натуральне число степенем 2
+ RecursiveCheckPowerOfTwo(number int) -> bool
```go
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
```
---
### Task 13
Реалізовано рекурсивні функції знаходження суми перших n парних та непарних невід'ємних цілих чисел
+ RecursiveSumFirstEvenNumbers(n uint) -> uint
```go
func RecursiveSumFirstEvenNumbers(n uint) uint {
	if n == 0 {
		return 0
	}
	return 2*n + RecursiveSumFirstEvenNumbers(n-1)
}
```
+ RecursiveSumFirstOddNumbers(n uint) -> uint
```go
func RecursiveSumFirstOddNumbers(n uint) uint {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return 2*n - 1 + RecursiveSumFirstOddNumbers(n-1)
}
```
---
### Task 14
Реалізовано рекурсивну функцію для розв'язання задачі про "Ханойські вежі"
+ HanoiTowers(count int, from, to, auxiliary string) -> int
```go
func HanoiTowers(count int, from, to, auxiliary string) int {
	var iterationCount int
	hanoiTowers(count, from, to, auxiliary, &iterationCount)
	return iterationCount
}

func hanoiTowers(count int, from, to, auxiliary string, iterationCounter *int) {
	*iterationCounter += 1

	if count == 1 {
		fmt.Printf("Disk %d from %s to %s\n", count, from, to)
		return
	}

	hanoiTowers(count-1, from, auxiliary, to, iterationCounter)
	fmt.Printf("Disk %d from %s to %s\n", count, from, to)
	hanoiTowers(count-1, auxiliary, to, from, iterationCounter)
}
```
+ HanoiTowersMinimumSteps(diskCount uint) -> uint
```go
func HanoiTowersMinimumSteps(diskCount uint) uint {
	if diskCount < 1 || diskCount > 63 {
		return 0
	}
	return uint(math.Pow(2, float64(diskCount))) - 1
}
```
---