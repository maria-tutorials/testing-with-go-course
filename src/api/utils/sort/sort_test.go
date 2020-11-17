package sort

import "testing"

//whitebox testing because it's inside the same package

func TestBubbleSort(t *testing.T) {
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	BubbleSort(elements)

	if elements[0] != 0 {
		t.Error("Expected first element to be 0. Got", elements[0])
	}

	if elements[len(elements)-1] != 9 {
		t.Error("Expected last element to be 9. Got", elements[len(elements)-1])
	}
}

func TestGoodSort(t *testing.T) {
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	GoodSort(elements)

	if elements[0] != 0 {
		t.Error("Expected first element to be 0. Got", elements[0])
	}

	if elements[len(elements)-1] != 9 {
		t.Error("Expected last element to be 9. Got", elements[len(elements)-1])
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

func BenchmarkBubbleSort(b *testing.B) {
	elements := generateElements(10000)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkGoodSort(b *testing.B) {
	elements := generateElements(10000)
	for i := 0; i < b.N; i++ {
		GoodSort(elements)
	}
}
