package collections

func SliceAtIndex[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func SliceOutwards[T any](slice []T, from int, to int) []T {
	return append(slice[:from], slice[to+1:]...)
}

func MinMaxInt(slice []int) (int, int) {
	if slice == nil || len(slice) == 0 {
		panic("Slice cannot be empty!")
	}

	var min = slice[0]
	var max = slice[0]
	for _, item := range slice {
		if min > item {
			min = item
		}

		if max < item {
			max = item
		}
	}

	return min, max
}
