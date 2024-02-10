package sort

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

func BubbleSortStat(n uint) (uint, uint) {
	C := ((n - 1) * n) / 2
	M := ((n - 1) * n) / 4
	return C, M
}
