package bubble

import (
	"sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	ints := []int{7, 3, 4}
	checkBubbleSort(t, ints)

	ints = []int{45, 56, 67, 23, 45, 6}
	checkBubbleSort(t, ints)

	ints = []int{7, 45, 23, 67, 87}
	checkBubbleSort(t, ints)

	ints = []int{7, 3, 4, 5, 6, 1}
	checkBubbleSort(t, ints)
}

func checkBubbleSort(t *testing.T, ints []int) {
	t.Helper()
	BubbleSort(ints)
	if !sort.IntsAreSorted(ints) {
		t.Errorf("Ints not sorted")
	}
}
