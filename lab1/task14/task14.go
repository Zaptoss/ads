package task14

import (
	"fmt"
	"math"
)

// Реалізувати алгоритм для розв’язання задачі «Ханойські вежі».
// Виписати послідовність ходів для перекладання n дисків вежі (n = 2; 3; 4; 5 дисків, використати онлайн гру).

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

func HanoiTowersMinimumSteps(diskCount uint) uint {
	if diskCount < 1 || diskCount > 63 {
		return 0
	}
	return uint(math.Pow(2, float64(diskCount))) - 1
}
