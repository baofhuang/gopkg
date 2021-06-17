package maths

//单元测试文件
import "testing"

//加法
func Test_Add(t *testing.T) {
	var (
		expected int
		actual   int
	)
	actual, _ = Add(1, 2)
	expected = 3
	if actual != expected {
		t.Errorf("Message:expected:%d\tactual:%d", expected, actual)
	}
}

//减法
func Test_Sub(t *testing.T) {
	var (
		expected int
		actual   int
	)
	actual, _ = Sub(1, 2)
	expected = -1
	if actual != expected {
		t.Errorf("Message:expected:%d\tactual:%d", expected, actual)
	}
}

//乘法
func Test_Mutil(t *testing.T) {
	var (
		expected float32
		actual   float32
	)
	actual, _ = Mutil(1, 2)
	expected = 2
	if actual != expected {
		t.Errorf("Message:expected:%v\tactual:%v", expected, actual)
	}
}

//乘法
func Test_Div(t *testing.T) {
	var (
		expected float32
		actual   float32
	)
	actual, _ = Div(1, 2)
	expected = 0.5
	if actual != expected {
		t.Errorf("Message:expected:%f\tactual:%f", expected, actual)
	}
}

//模运算
func Test_Mod(t *testing.T) {
	var (
		expected int
		actual   int
	)
	actual, _ = Mod(1, 2)
	expected = 1
	if actual != expected {
		t.Errorf("Message:expected:%d\tactual:%d", expected, actual)
	}
}

//不定个数参数运算
func Test_SumAllInt(t *testing.T) {
	var (
		expected int
		actual   int
	)
	actual, _ = SumAllInt(1, 2, 3, 4, 5)
	expected = 15
	if actual != expected {
		t.Errorf("Message:expected:%d\tactual:%d", expected, actual)
	}
}

//参数为slice
func Test_SumSliceInt(t *testing.T) {
	var (
		expected int
		actual   int
	)
	actual, _ = SumSliceInt([]int{1, 2, 3, 4, 5})
	expected = 15
	if actual != expected {
		t.Errorf("Message:expected:%d\tactual:%d", expected, actual)
	}
}
