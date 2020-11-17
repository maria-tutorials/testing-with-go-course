package services

import "../utils/sort"

//import "https://github.com/maria-tutorials/testing-with-go-course/src/api/utils/sort"

func Sort(elements []int) {
	if len(elements) >= 10000 {
		sort.GoodSort(elements)
		return
	}
	sort.BubbleSort(elements)
}
