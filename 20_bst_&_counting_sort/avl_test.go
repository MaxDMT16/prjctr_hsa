package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/emirpasic/gods/trees/avltree"
)

type keyGeneratorFunc func() int

func insert(b *testing.B, tree *avltree.Tree, n int, generateKey keyGeneratorFunc) {
	for j := 0; j < n; j++ {
		tree.Put(generateKey(), 1)
	}
}

var table = []int{10, 100, 1000, 10_000, 100_000, 1_000_000}

func BenchmarkInsert_DifferentNumbers(b *testing.B) {
	tree := avltree.NewWithIntComparator()

	for _, v := range table {
		b.Run(fmt.Sprintf("insert_%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				insert(b, tree, v, func() int { return rand.Int() })

				tree.Clear()
			}
		})
	}
}

func BenchmarkInsert_SameNumber(b *testing.B) {
	tree := avltree.NewWithIntComparator()

	for _, v := range table {
		b.Run(fmt.Sprintf("insert_%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				insert(b, tree, v, func() int { return 1 })

				tree.Clear()
			}
		})
	}
}

func BenchmarkFind_SameNumber(b *testing.B) {
	tree := avltree.NewWithIntComparator()

	for _, v := range table {
		b.Run(fmt.Sprintf("find_%d", v), func(b *testing.B) {
			insert(b, tree, v, func() int { return 1 })

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tree.Get(1)
			}
		})
	}
}

func BenchmarkFind_DifferentNumber(b *testing.B) {
	tree := avltree.NewWithIntComparator()

	for _, v := range table {
		b.Run(fmt.Sprintf("find_%d", v), func(b *testing.B) {
			insert(b, tree, v, func() int { return rand.Int() })

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tree.Get(i)
			}
		})
	}
}

func BenchmarkRemove_DifferentNumber(b *testing.B) {
	tree := avltree.NewWithIntComparator()

	for _, v := range table {
		b.Run(fmt.Sprintf("remove_%d", v), func(b *testing.B) {
			insert(b, tree, v, func() int { return rand.Int() })

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tree.Remove(i)
			}
		})
	}
}

func BenchmarkRemove_SameNumber(b *testing.B) {
	tree := avltree.NewWithIntComparator()

	for _, v := range table {
		b.Run(fmt.Sprintf("remove_%d", v), func(b *testing.B) {
			insert(b, tree, v, func() int { return 1 })

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tree.Remove(i)
			}
		})
	}
}
