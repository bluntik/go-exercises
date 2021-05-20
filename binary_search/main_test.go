package main

import (
	"math/rand"
	"testing"
	"time"
)

type dataProvider struct {
	haystack []int
	needle   int
	expected bool
}

var tests = []dataProvider{
	{[]int{}, 0, false},
	{[]int{1, 2}, 1, true},
	{[]int{1, 2, 3, 4, 5, 6}, 2, true},
	{[]int{1, 2, 3, 4, 5, 6}, 1, true},
	{[]int{1, 2, 3, 4, 5, 6}, 6, true},
	{[]int{1, 2, 3, 4, 5, 6}, 7, false},
}

func TestSearch(t *testing.T) {
	for _, test := range tests {
		found, _ := Search(test.haystack, test.needle)

		if test.expected != found {
			t.Error(
				"For haystack", test.haystack,
				"Needle", test.needle,
				"got", found,
			)
		}
	}
}

func BenchmarkSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Search([]int{1, 2}, 1)
	}
}

func BenchmarkSearch1Mln(b *testing.B) {
	length := 1_000_000

	haystack := make([]int, length+1)

	for i := 0; i <= length; i++ {
		haystack[i] = i
	}

	needle := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(haystack))

	for i := 0; i < b.N; i++ {
		Search(haystack, needle)
	}
}
