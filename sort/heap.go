package main

func HeapSort(arr []int) {
	N := len(arr)
	for start := N / 2; start > -1; start-- {
		maxHeapify(arr, start, N-1)
	}
	for end := N - 1; end > 0; end-- {
		arr[end], arr[0] = arr[0], arr[end]
		maxHeapify(arr, 0, end-1)
	}
}

func maxHeapify(arr []int, start, end int) {
	root := start
	for true {
		child := root*2 + 1
		if child > end {
			break
		}
		if child+1 < end && arr[child] < arr[child+1] {
			child = child + 1
		}
		if arr[root] < arr[child] {
			arr[root], arr[child] = arr[child], arr[root]
			root = child
		} else {
			break
		}
	}
}
