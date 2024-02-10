package main

import (
	"fmt"
	"lab2/sort"
	"math/rand"
	"time"
	// "time"
)

func timeCatch(my_func func([]int) (uint, uint), arr []int) (uint, uint, time.Duration) {
	repeatCount := 100
	var avgDuration time.Duration
	var avgC, avgM, C, M uint

	for i := 0; i < repeatCount; i++ {
		rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
		start := time.Now()
		C, M = my_func(arr)
		avgDuration += time.Since(start)
		avgC += C
		avgM += M
	}
	// fmt.Printf("Used time: %d\n", duration)
	return avgC / uint(repeatCount), avgM / uint(repeatCount), avgDuration / time.Duration(repeatCount)
}

func testSort(my_func func([]int) (uint, uint), statFunc func(uint) (uint, uint), sortName string, arraySize int) {
	var C, M uint
	var duration time.Duration

	fmt.Printf("==>>\t%sSort\t| n = %d\n", sortName, arraySize)

	C, M = statFunc(uint(arraySize))
	fmt.Printf("|- %s sort theoretical statistic: \tC = %d; \tM = %d\n", sortName, C, M)

	array := rand.Perm(arraySize)

	C, M, duration = timeCatch(my_func, array)
	fmt.Printf("|- %s sort experimental statistic: \tC = %d; \tM = %d\n|- Used time:\t%s\n\n", sortName, C, M, duration)
}

func main() {

	testCounts := []int{100, 1000, 10000}

	for _, count := range testCounts {
		testSort(sort.QuickSort, sort.QuickSortStat, "Quick", count)
		testSort(sort.SelectionSort, sort.SelectionSortStat, "Selection", count)
	}

}
