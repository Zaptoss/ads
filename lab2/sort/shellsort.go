package sort

import (
	"math"
)

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

func ShellSortStat(n uint) (uint, uint) {
	C := uint(math.Pow(float64(n), 1.2))
	M := C
	return C, M
}
