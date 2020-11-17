package services

import "testing"

func TestSort(t *testing.T) {
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	Sort(elements)

	if elements[0] != 0 {
		t.Error("Expected first element to be 0. Got", elements[0])
	}

	if elements[len(elements)-1] != 9 {
		t.Error("Expected last element to be 9. Got", elements[len(elements)-1])
	}
}
