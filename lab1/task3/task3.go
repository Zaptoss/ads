package task3

// Числа Фібоначчі f(n) обчислюються за формулами f(0) = f(1) = 1;
// f(n) = f(n-1) + f(n-2) при n=2,3,... Реалізувати функцію, яка за заданим номером n
// обчислюватиме значення f(n).

func RecursiveFibonachi(number uint) uint {
	if number <= 1 {
		return 1
	}
	return RecursiveFibonachi(number-1) + RecursiveFibonachi(number-2)
}
