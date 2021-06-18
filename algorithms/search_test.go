package algorithms

import (
	"testing"
)

//直接查找
func Test_DirectSearch(t *testing.T) {
	actual, err := DirectSearch([]int{1, 3, 5, 7, 9, 2, 4, 6}, 15)
	expected := -1
	if expected != actual {
		t.Errorf("Message:expected:%d actual:%d", expected, actual)
		t.Errorf("Error:%v", err)
	}
}

//二分查找
func Test_BinarySearch(t *testing.T) {
	actual, err := BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5)
	expected := 4
	if expected != actual {
		t.Errorf("Message:expected:%d actual:%d", expected, actual)
		t.Errorf("Error:%v", err)
	}
}
