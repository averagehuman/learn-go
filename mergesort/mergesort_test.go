package mergesort

import "fmt"

func ExampleMergeSort() {
	fmt.Println(MergeSort([]int64{1}))
	fmt.Println(MergeSort([]int64{2, 1}))
	fmt.Println(MergeSort([]int64{1, 2}))
	fmt.Println(MergeSort([]int64{2, 3, 9, 10, 456, 5, 4, 33, 39, 25, 1, 2}))
	// Output:
	// [1]
	// [1 2]
	// [1 2]
	// [1 2 2 3 4 5 9 10 25 33 39 456]
}
