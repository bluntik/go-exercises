package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	defer func(start time.Time) {
		elapsed := time.Since(start)

		fmt.Printf("execution time %v \n", elapsed)
	}(time.Now())

	length := 1_000_000

	haystack := make([]int, length+1)

	for i := 0; i <= length; i++ {
		haystack[i] = i
	}

	needle := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(haystack))

	if found, index := Search(haystack, needle); found {
		fmt.Printf("Found value `%d` with index %d.\n", needle, index)
	} else {
		fmt.Printf("Value `%d` was NOT found.\n", needle)
	}
}

func Search(haystack []int, needle int) (bool, int) {
	low := 0
	high := len(haystack) - 1

	for low <= high {
		mid := (low + high) / 2

		guess := haystack[mid]

		if guess == needle {
			return true, mid
		} else if guess >= needle {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return false, 0
}
