package main

import "testing"

//测试带锁并发读
func BenchmarkMutexMap_ConcurrentReadConcurrentRead(b *testing.B) {
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		mmp := MutexMap{m: make(map[string]string)}
		mmp.ConcurrentRead()
	}
}

//测试不带锁并发读
func BenchmarkMutexMap_ConcurrentReadConcurrentReadLock(b *testing.B) {
	b.ResetTimer()
	for i:=0;i<b.N;i++ {
		mmp := MutexMap{m: make(map[string]string)}
		mmp.ConcurrentReadLock()
	}
}

//测试并发写
func BenchmarkMutexMap_ConcurrentWrite(b *testing.B) {
	b.ResetTimer()
	for i:=0;i<b.N;i++ {
		mmp := MutexMap{m: make(map[string]string)}
		mmp.ConcurrentWrite()
	}
}

//测试并发读写
func BenchmarkMutexMap_ReadAndWrite(b *testing.B) {
	b.ResetTimer()
	for i:=0;i<b.N;i++ {
		mmp := MutexMap{m: make(map[string]string)}
		mmp.ReadAndWrite()
	}
}
