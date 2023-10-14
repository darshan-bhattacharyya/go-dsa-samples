package main

func MergeSort(array []interface{}, left int, right int, compare func(i, j int) bool) []interface{} {
	if left == right {
		return array
	} else {
		mid := (left + right) / 2

		MergeSort(array, left, mid, compare)
		MergeSort(array, mid+1, right, compare)

		merge(array, left, mid, right, compare)
	}
	return array
}

func merge(array []interface{}, left, mid, right int, compare func(i, j int) bool) {
	i := left
	j := mid + 1
	sortedArray := make([]interface{}, 0)
	for i <= mid && j <= right {
		if compare(i, j) {
			sortedArray = append(sortedArray, array[i])
			i += 1
		} else {
			sortedArray = append(sortedArray, array[j])
			j += 1
		}
	}

	for i <= mid {
		sortedArray = append(sortedArray, array[i])
		i += 1
	}

	for j <= right {
		sortedArray = append(sortedArray, array[j])
		j += 1
	}

	for e, k := 0, left; k <= right; e, k = e+1, k+1 {
		array[k] = sortedArray[e]
	}
}

func QuickSort(array []interface{}, left int, right int, compare func(i, j int) bool) []interface{} {
	if left < right {
		pivot := partition(array, left, right, compare)

		QuickSort(array, left, pivot-1, compare)
		QuickSort(array, pivot+1, right, compare)
	}
	return array
}

func partition(array []interface{}, low int, high int, compare func(i, j int) bool) int {
	i := low - 1
	for j := low; j < high; j += 1 {
		if compare(j, high) {
			i += 1
			array[i], array[j] = array[j], array[i]
		}
	}

	array[i+1], array[high] = array[high], array[i+1]

	return i + 1
}
