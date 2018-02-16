package mergesort

func MergeSort(seq []int64) []int64 {
	d := len(seq)
	if d < 2 {
		return seq
	}
	left := seq[:d/2]
	right := seq[d/2:]

	return merge(MergeSort(left), MergeSort(right))
}

func merge(left []int64, right []int64) []int64 {
	sorted := []int64{}

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			sorted = append(sorted, left[0])

}
