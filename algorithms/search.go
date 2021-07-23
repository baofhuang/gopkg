package algorithms

import (
	"errors"
)

//直接查找
func DirectSearch(data []int, key int) (int, error) {
	for k, value := range data {
		if value == key {
			return k, nil
		}
	}
	return -1, errors.New("未找到")
}

//折半查找，针对有序序列
func BinarySearch(data []int, key int) (int, error) {
	if !IsInOrder(data) {
		return -1, errors.New("输入数据应为有序数据")
	}
	//调用排序算法判断序列是否为有序
	return binSearch(data, key, 0, len(data)-1)
}
func binSearch(data []int, key int, start int, end int) (int, error) { //仅由包内部调用
	if start > end {
		return -1, errors.New("未找到")
	}
	mid := (start + end) / 2
	if data[mid] == key {
		return mid, nil
	}
	if data[mid] < key {
		binSearch(data, key, start, mid-1)
	} else {
		binSearch(data, key, mid+1, end)
	}
	return -1, errors.New("未找到")
}
