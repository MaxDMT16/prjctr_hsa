package main

import (
	"fmt"

	avl "github.com/emirpasic/gods/trees/avltree"
)

func main() {
	avlDemo()
}

func avlDemo() {
	tree := avl.NewWithIntComparator() // empty(keys are of type int)

	tree.Put(1, "x") // 1->x
	tree.Put(2, "b") // 1->x, 2->b (in order)
	tree.Put(1, "a") // 1->a, 2->b (in order, replacement)
	tree.Put(3, "c") // 1->a, 2->b, 3->c (in order)
	tree.Put(4, "d") // 1->a, 2->b, 3->c, 4->d (in order)
	tree.Put(5, "e") // 1->a, 2->b, 3->c, 4->d, 5->e (in order)
	tree.Put(6, "f") // 1->a, 2->b, 3->c, 4->d, 5->e, 6->f (in order)

	n := tree.GetNode(1)
	fmt.Println(n.Value)

	_ = tree.Values() // []interface {}{"a", "b", "c", "d", "e", "f"} (in order)
	_ = tree.Keys()   // []interface {}{1, 2, 3, 4, 5, 6} (in order)

	tree.Remove(2) // 1->a, 3->c, 4->d, 5->e, 6->f (in order)
	fmt.Println(tree)
	//
	//  AVLTree
	//  │       ┌── 6
	//  │   ┌── 5
	//  └── 4
	//      └── 3
	//          └── 1

	tree.Clear() // empty
	tree.Empty() // true
	tree.Size()  // 0
}

func countingSortDemo() {
	slice := []int{7, 1, 5, 2, 2}
	fmt.Println(CountingSort(slice)) // [0 1 2 2 5 7]

	slice1 := []int{1, 2, 3, 6, 4, 5, 4, 6, 7, 8}
	fmt.Println(CountingSort(slice1)) // [0 1 2 3 4 4 5 6 6 7 8]
	fmt.Println(slice1)               // [1 2 3 6 4 5 4 6 7 8]

	slice2 := []int{20, 370, 45, 75, 410, 1802, 24, 2, 66}
	fmt.Println(CountingSort(slice2))
	// [0 2 20 24 45 66 75 370 410 1802]

	fmt.Println(slice2)
	// [20 370 45 75 410 1802 24 2 66]
}
