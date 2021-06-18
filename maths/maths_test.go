package maths

//单元测试文件
import "testing"

//加法
func Test_Add(t *testing.T) {
	actual, err := Add(1, 2)
	expected = 3
	if actual != expected {
		t.Errorf("Message:expected:%d\tactual:%d", expected, actual)
		t.Errorf("Error:%v", err)
	}
}

//减法
func Test_Sub(t *testing.T) {
	actual, err := Sub(1, 2)
	expected := -1
	if actual != expected {
		t.Errorf("Message:expected:%d\tactual:%d", expected, actual)
		t.Errorf("Error:%v", err)
	}
}

//乘法
func Test_Multi(t *testing.T) {
	actual, err := Multi(1, 2)
	expected := 2
	if actual != expected {
		t.Errorf("Message:expected:%v\tactual:%v", expected, actual)
		t.Errorf("Error:%v", err)
	}
}

//除法
func Test_Div(t *testing.T) {
	actual, err := Div(1, 2)
	expected := 0.5
	if actual != expected {
		t.Errorf("Message:expected:%f\tactual:%f", expected, actual)
		t.Errorf("Error:%v", err)
	}
}

//模运算
func Test_Mod(t *testing.T) {
	actual, err := Mod(1, 2)
	expected = 1
	if actual != expected {
		t.Errorf("Message:expected:%d\tactual:%d", expected, actual)
		t.Errorf("Error:%v", err)
	}
}

//指数运算
func Test_Power(t *testing.T) {
	var (
		expected float32
		actual   float32
	)
	actual, _ = Power(2, -2)
	expected = 0.25

	if actual != expected {
		t.Errorf("Message:expected:%f\tactual:%f", expected, actual)
	}
}

//不定个数参数运算
func Test_SumAllInt(t *testing.T) {
	actual, err := SumAllInt(1, 2, 3, 4, 5)
	expected := 15
	if actual != expected {
		t.Errorf("Message:expected:%d\tactual:%d", expected, actual)
		t.Errorf("Error:%v", err)
	}
}

//参数为slice
func Test_SumSliceInt(t *testing.T) {
	actual, err := SumSliceInt([]int{1, 2, 3, 4, 5})
	expected := 15
	if actual != expected {
		t.Errorf("Message:expected:%d\tactual:%d", expected, actual)
		t.Errorf("Error:%v", err)
	}
}
