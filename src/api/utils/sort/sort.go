package sort

// BubbleSort will sort a list of integers using the bubble sort algorithm.
func BubbleSort(list []int) {
	len := len(list)
	swapped := false

	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
