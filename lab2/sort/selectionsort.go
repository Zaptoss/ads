package sort

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

func SelectionSortStat(n uint) (uint, uint) {
	C := ((n - 1) * n) / 2
	M := (n - 1) / 2
	return C, M
}
