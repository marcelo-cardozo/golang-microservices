package utils

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{5, 1, 7, 3, 9}

	BubbleSort(arr)

	assert.Equal(t, 5, len(arr))

	assert.Equal(t, 1, arr[0])
	assert.Equal(t, 3, arr[1])
	assert.Equal(t, 5, arr[2])
	assert.Equal(t, 7, arr[3])
	assert.Equal(t, 9, arr[4])
}

func getElements(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Int()
	}
	return arr
}

func BenchmarkBubbleSort100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := getElements(100)

		BubbleSort(arr)
	}
}
func BenchmarkSort100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := getElements(100)

		sort.Ints(arr)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := getElements(1000)

		BubbleSort(arr)
	}
}
func BenchmarkSort1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := getElements(1000)

		sort.Ints(arr)
	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := getElements(100000)

		BubbleSort(arr)
	}
}
func BenchmarkSort100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := getElements(100000)

		sort.Ints(arr)
	}
}
