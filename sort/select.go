package main

func SelectSort(arr []int) {
	len := len(arr)
	for i := 0; i < len-1; i++ {
		min := i
		for j := i + 1; j < len; j++ {
			if arr[j] < arr[i] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}
