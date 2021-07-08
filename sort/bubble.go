package main

func main() {

}

func BubbleSort(arr []int) {
	swapped := false
	len := len(arr)
	for swapped {
		swapped = false
		for i := 0; i < len-1; i++ {
			if arr[i] < arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
			swapped = true
		}
	}
}
