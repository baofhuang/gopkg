package algorithms

//检查是否为有序序列
func IsInOrder(data []int) bool {
	if len(data) == 0 || len(data) == 1 {
		return true
	}
	if data[0] < data[1] {
		for i := 1; i < len(data); i++ {
			if data[i] < data[i-1] {
				return false
			}
		}
	} else {
		for i := 1; i < len(data); i++ {
			if data[i] > data[i-1] {
				return false
			}
		}
	}
	return true
}

//插入排序
func InsertSort(data []int) bool {
	if IsInOrder(data) {
		return false
	}
	var mid = 1
	for ; mid < len(data); mid++ {
		value := data[mid]
		for i := mid - 1; i > 0; i-- {
			if data[i] < value { //降序排序
				data[mid] = data[i]
				data[i] = value
				break
			}
			// if data[i] > value { //升序排序
			// 	data[mid] = data[i]
			// 	data[i] = value
			// 	break
			// }
		}

	}
	return true
}

//选择排序

//冒泡排序

//快排

//合并排序

//堆排序 - 建堆 - 堆调整
