package services

import "testing"

func TestSort(t *testing.T) {
	elements := generateElements(10)
	Sort(elements)

	if elements[0] != 0 {
		t.Error("Expected first element to be 0. Got", elements[0])
	}

	if elements[len(elements)-1] != 9 {
		t.Error("Expected last element to be 9. Got", elements[len(elements)-1])
	}
}

func TestBIGSort(t *testing.T) {
	elements := generateElements(10000)
	Sort(elements)

	if elements[0] != 0 {
		t.Error("Expected first element to be 0. Got", elements[0])
	}

	if elements[len(elements)-1] != 9999 {
		t.Error("Expected last element to be 9999. Got", elements[len(elements)-1])
	}
}

//helper
func generateElements(n int) []int {
	result := make([]int, n)
	j := 0
	for i := n - 1; i > 0; i-- {
		result[j] = i
		j++
	}
	return result
}
