package mergesort

func MergeSort(seq []int64) []int64 {
	// Recursively sort each half of seq and merge the result
	d := len(seq)
	if d < 2 {
		return seq
	}
	left := seq[:d/2]
	right := seq[d/2:]

	return merge(MergeSort(left), MergeSort(right))
}

func merge(left []int64, right []int64) []int64 {
	merged := []int64{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}
	if i < len(left) {
		merged = append(merged, left[i:]...)
	}
	if j < len(right) {
		merged = append(merged, right[j:]...)
	}
	return merged
}
