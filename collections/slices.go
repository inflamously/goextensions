package collections

func SliceAtIndex[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func SliceOutwards[T any](slice []T, from int, to int) []T {
	return append(slice[:from], slice[to+1:]...)
}

/*
Reverse slice inline
*/
func Reverse[T any](slice []T) {
	if len(slice) < 2 {
		return
	}

	var tmp T
	for i, j := 0, len(slice)-1; i < len(slice); i, j = i+1, j-1 {
		if i > j {
			break
		}

		tmp = slice[i]
		slice[i] = slice[j]
		slice[j] = tmp
	}
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
