package sort

import "math"

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

func QuickSortStat(n uint) (uint, uint) {
	C := n * uint(math.Log2(float64(n)))
	M := C
	return C, M
}
