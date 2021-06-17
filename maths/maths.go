package maths

import "errors"
//加法
func Add(num1, num2 int) (int, error) {
	return num1 + num2, nil
}
//减法
func Sub(num1, num2 int) (int, error) {
	return num1 + num2, nil
}

//乘法
func Mutil(num1, num2 float32) (float32, error) {
	return num1 + num2, nil
}

//乘法
func Div(num1, num2 float32) (float32, error) {
	if num2 == 0 {
		return 0, errors.New("被除数为0！")
	}
	return num1 / num2, nil
}

//模运算
func Mod(num1, num2 int) (int, error) {
	return num1 % num2, nil
}
