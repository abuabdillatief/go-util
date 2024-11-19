package slices

func GetMax[T ~int](slice []T) T {
	max := slice[0]
	for _, v := range slice {
		if v > max {
			max = v
		}
	}

	return max
}

func GetMin[T ~int](slice []T) T {
	max := slice[0]
	for _, v := range slice {
		if v < max {
			max = v
		}
	}

	return max
}

func GetSum[T ~int](slice []T) T {
	sum := slice[0]
	for _, v := range slice[1:] {
		sum += v
	}

	return sum
}

func GetAverage[T ~int](slice []T) float64 {
	sum := GetSum(slice)
	return float64(sum) / float64(len(slice))
}

func GetMedian[T ~int](slice []T) float64 {
	sorted := make([]T, len(slice))
	copy(sorted, slice)
	QuickSort(sorted)

	mid := len(sorted) / 2
	if len(sorted)%2 == 0 {
		return float64(sorted[mid-1]+sorted[mid]) / 2
	}

	return float64(sorted[mid])
}

func QuickSort[T ~int](slice []T) {
	if len(slice) < 2 {
		return
	}

	left, right := 0, len(slice)-1
	pivot := len(slice) / 2

	slice[pivot], slice[right] = slice[right], slice[pivot]

	for i := range slice {
		if slice[i] < slice[right] {
			slice[i], slice[left] = slice[left], slice[i]
			left++
		}
	}

	slice[left], slice[right] = slice[right], slice[left]

	QuickSort(slice[:left])
	QuickSort(slice[left+1:])
}