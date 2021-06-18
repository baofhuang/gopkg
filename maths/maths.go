package maths

import (
	"errors"
)

//加法
func Add(num1, num2 int) (int, error) {
	return num1 + num2, nil
}

//减法
func Sub(num1, num2 int) (int, error) {
	return num1 - num2, nil
}

//乘法
func Multi(num1, num2 float32) (float32, error) {
	return num1 * num2, nil
}

//除法
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

//指数运算
func Power(base float32, exponent int) (float32, error) {
	var sum = base
	if exponent < 0 {
		for exponent++; exponent < 0; exponent++ {
			sum *= base
		}
		sum = 1.0 / sum
	} else {
		for exponent--; exponent > 0; exponent-- {
			sum *= base
		}
	}
	return sum, nil
}

//不定个数参数运算
func SumAllInt(args ...int) (int, int) {
	var sum = 0
	var count = 0
	for _, value := range args {
		sum += value
		count++
	}
	return sum, count
}

//参数为slice
func SumSliceInt(args []int) (int, int) {
	var sum = 0
	var count = 0
	for _, value := range args {
		sum += value
		count++
	}
	return sum, count
}
