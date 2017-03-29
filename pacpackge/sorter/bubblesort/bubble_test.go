package bubblesort

import "testing"

func TestBubbleSort1(t *testing.T) {
	values := []int{2, 4, 73, 3, 9}
	BubbleSort(values)
	if values[0] != 2 || values[1] != 3 || values[2] != 4 || values[3] != 9 || values[4] != 73 {
		t.Error("BubbleSort() failed. Got", values, "Expexted 2,3,4,9,73")
	}
}
