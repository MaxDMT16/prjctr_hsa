package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkCountingSort(b *testing.B) {
	table := []int{1, 10, 100, 1000, 10_000, 100_000, 1_000_000}
	b.ResetTimer()

	for _, v := range table {
		b.Run(fmt.Sprintf("counting_sort_%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				input := rand.Perm(v)

				CountingSort(input)
			}
		})
	}
}

func BenchmarkCountingSort_MyInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountingSort([]int{0, 3_000_000_000, 3})
	}
}
