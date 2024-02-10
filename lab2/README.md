# LAB2 ADS
#### В даній лабораторній роботі було використано мову програмування GoLang
---
#### Реалізовано 5 алгоритмів сортування масиву:
+ BubbleSort(arr []int) -> (C, M uint)
```go
func BubbleSort(arr []int) (uint, uint) {
	var C, M uint
	bubbleSort(arr, &C, &M)

	return C, M
}

func bubbleSort(arr []int, C, M *uint) {
	arrLength := len(arr)
	var move bool
	for i := arrLength - 1; i > 0; i-- {
		move = false
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				move = true
				*M++
			}
			*C++
		}

		if !move {
			break
		}
	}

}
```
+ SelectionSort(arr []int) -> (C, M uint)
```go
func SelectionSort(arr []int) (uint, uint) {
	var C, M uint
	selectionSort(arr, &C, &M)

	return C, M
}

func selectionSort(arr []int, C, M *uint) {
	arrLength := len(arr)
	var minItem [2]int

	for i := 0; i < arrLength-1; i++ {
		minItem[0] = i
		minItem[1] = arr[i]

		for j := i + 1; j < arrLength; j++ {
			if arr[j] < minItem[1] {
				minItem[1] = arr[j]
				minItem[0] = j
			}
			*C++
		}

		if minItem[0] != i {
			arr[minItem[0]], arr[i] = arr[i], minItem[1]
			*M++
		}
	}
}
```
+ InsertionSort(arr []int) -> (C, M uint)
```go
func InsertionSort(arr []int) (uint, uint) {
	var C, M uint
	insertionSort(arr, &C, &M)
	return C, M
}

func insertionSort(arr []int, C, M *uint) {
	var key int

	for i := 1; i < len(arr); i++ {
		key = arr[i]
		j := i - 1

		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j--
			*M++
			*C++
		}

		arr[j+1] = key
		*M++
	}
}
```
+ QuickSort(arr []int) -> (C, M uint)
```go
func QuickSort(arr []int) (uint, uint) {
	var C, M uint
	quickSort(arr, &C, &M)

	return C, M
}

func quickSort(arr []int, C, M *uint) {
	arrLength := len(arr)

	if arrLength <= 1 {
		return
	}

	base := arr[0]
	basePointer := 1

	for i := 1; i < arrLength; i++ {
		if arr[i] <= base {
			arr[basePointer], arr[i] = arr[i], arr[basePointer]
			basePointer++
			*M++
		}
		*C++
	}

	arr[basePointer-1], arr[0] = arr[0], arr[basePointer-1]
	*M++

	quickSort(arr[:basePointer-1], C, M)
	quickSort(arr[basePointer:], C, M)
}
```
+ ShellSort(arr []int) -> (C, M uint)
```go
func ShellSort(arr []int) (uint, uint) {
	var C, M uint
	shellSort(arr, &C, &M)

	return C, M
}

func shellSort(arr []int, C, M *uint) {
	K := int(math.Log2(float64(len(arr)))) - 1

	distance := make([]int, K)

	distance[K-1] = 1
	for i := K - 1; i > 0; i-- {
		distance[i-1] = distance[i]*2 + 1
	}

	for _, i := range distance {
		for j := 0; j < i; j++ {
			insertionStepSort(arr[j:], i, C, M)
		}
	}

}

func insertionStepSort(arr []int, step int, C, M *uint) {
	var key int

	for i := step; i < len(arr); i += step {
		key = arr[i]
		j := i - step

		for j >= 0 && key < arr[j] {
			arr[j+step] = arr[j]
			j -= step
			*C++
			*M++
		}

		arr[j+step] = key
		*M++
	}
}
```

В індивідуальному варіанті необхідно порівняти алгоритми Selection sort і Quick sort, порахувати час виконання сортування і кількість переміщень та порівнянь. Для усереднених результатів було проведено по декілька тестів для кожного із запропанованих розмірів масиву.

Отримані дані узагальнено в таблиці:

<table>
        <tr>
            <th></th>
            <th colspan=5>Selection Sort</th>
            <th colspan=5>Quick Sort</th>
        </tr>
        <tr>
            <th rowspan=2>N</th>
            <th colspan=2>К-ть копіювань (M)</th>
            <th colspan=2>К-ть порівнянь (C)</th>
            <th rowspan=2>Час (T)</th>
            <th colspan=2>К-ть копіювань (M)</th>
            <th colspan=2>К-ть порівнянь (C)</th>
            <th rowspan=2>Час (T)</th>
        </tr>
        <tr>
            <th>Теорет.</th>
            <th>Експерим.</th>
            <th>Теорет.</th>
            <th>Експерим.</th>
            <th>Теорет.</th>
            <th>Експерим.</th>
            <th>Теорет.</th>
            <th>Експерим.</th>
        </tr>
        <tr>
            <td>100</td>
            <td>49</td>
            <td>94</td>
            <td>4950</td>
            <td>4950</td>
            <td>5.641µs</td>
            <td>600</td>
            <td>388</td>
            <td>600</td>
            <td>651</td>
            <td>3.58µs</td>
        </tr>
        <tr>
            <td>1000</td>
            <td>499</td>
            <td>992</td>
            <td>499500</td>
            <td>499500</td>
            <td>307.56µs</td>
            <td>9000</td>
            <td>6166</td>
            <td>9000</td>
            <td>11113</td>
            <td>36.156µs</td>
        </tr>
        <tr>
            <td>10000</td>
            <td>4999</td>
            <td>9990</td>
            <td>49995000</td>
            <td>49995000</td>
            <td>26.363231ms</td>
            <td>130000</td>
            <td>84810</td>
            <td>130000</td>
            <td>155687</td>
            <td>497.307µs</td>
        </tr>
</table>
