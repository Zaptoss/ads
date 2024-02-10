package sort

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

func InsertionSortStat(n uint) (uint, uint) {
	C := ((2 + n) * (n - 1)) / 4
	M := ((n - 1) * n) / 4
	return C, M
}
