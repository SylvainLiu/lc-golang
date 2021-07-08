package main

func InsertSort(arr []int) {
	l := len(arr)
	for i := 1; i < l; i++ {
		tmp := arr[i]
		for j := i - 1; j >= 0; j-- {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			} else {
				arr[j+1] = tmp
				break
			}
		}
	}
}
