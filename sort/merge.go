package main

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}

func merge(a, b []int) []int {
	al, bl := len(a), len(b)
	res := make([]int, al+bl)

	i, j, k := 0, 0, 0
	for i < al && j < bl {
		if a[i] < b[j] {
			res[k] = a[i]
			i++
		} else {
			res[k] = b[j]
			j++
		}
		k++
	}

	for i != al {
		res[k] = a[i]
		i++
		k++
	}
	for j != bl {
		res[k] = b[i]
		i++
		k++
	}
	return res
}
