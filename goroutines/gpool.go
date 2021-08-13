package main

// 定义任务
type Task struct {
	//goroutine执行对象为函数：函数参数类型与个数不确定
	Handler func(args ...interface{})
}

//定义任务池:容量Capacity，使用量
type Pool struct {
}
