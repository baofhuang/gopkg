package algorithms

import (
	"testing"
)

//检查是否为有序序列
func Test_IsInOrder(t *testing.T) {
	actual := IsInOrder([]int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	expected := false
	if actual != expected {
		t.Errorf("Message:expected:%v actual:%v", expected, actual)
	}
}

//插入排序
func Test_InsertSort(t *testing.T) {
	data := []int{1, 3, 5, 7, 9, 2, 4, 6, 8}
	_ = InsertSort(data)
	actual := InsertSort(data)
	expected := true
	if actual != expected {
		t.Errorf("Message:expected:%v actual:%v", expected, actual)
	}
}
