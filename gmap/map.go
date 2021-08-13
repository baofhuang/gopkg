package main

import (
	"sync"
)

var wg sync.WaitGroup

type MutexMap struct {
	sync.Mutex
	m map[string]string
}

type SyncMap struct {
	Sm sync.Map
}

// ReadAndWrite 并发读写
func (mMap *MutexMap) ReadAndWrite() {
	wg.Add(2)
	go mMap.WriteMap()
	go mMap.ReadMapLock()
	wg.Wait()
}

// ConcurrentWrite 并发写
func (mMap *MutexMap) ConcurrentWrite() {
	wg.Add(2)
	go mMap.WriteMap()
	go mMap.WriteMap()
	wg.Wait()
}

// ConcurrentReadLock 带锁并发读
func (mMap *MutexMap) ConcurrentReadLock() {
	wg.Add(2)
	go mMap.ReadMapLock()
	go mMap.ReadMapLock()
	wg.Wait()
}

// ConcurrentRead 不带锁并发读
func (mMap *MutexMap) ConcurrentRead() {
	wg.Add(2)
	go mMap.ReadMap()
	go mMap.ReadMap()
	wg.Wait()
}

//并发写
func (mMap *MutexMap) WriteMap() {
	for i := 0; i < 1; i++ {
		mMap.Lock()
		mMap.m["test"] = "newValue"
		mMap.Unlock()
	}
	wg.Done()
}

//不带锁并发读
func (mMap *MutexMap) ReadMap() {
	for i := 0; i < 5000000; i++ {
		_ = mMap.m["test"]
	}
	wg.Done()
}

//带锁并发读
func (mMap *MutexMap) ReadMapLock() {
	for i := 0; i < 5000000; i++ {
		mMap.Lock()
		_ = mMap.m["test"]
		mMap.Unlock()
	}
	wg.Done()
}

// SyncMapRead 使用SyncMap多次读取数据
func (sm *SyncMap) SyncMapRead() {
	for i := 0; i < 5000000; i++ {
		sm.Sm.Load("test")
	}
	wg.Done()
}

//SyncMapWrite 使用SyncMap多次写数据
func (sm *SyncMap) SyncMapWrite() {
	for i := 0; i < 1; i++ {
		sm.Sm.Store("test", "newValue")
	}
	wg.Done()
}

// SyncMapConcurrentRead 并发读
func (sm *SyncMap) SyncMapConcurrentRead() {
	wg.Add(2)
	go sm.SyncMapRead()
	go sm.SyncMapRead()
	wg.Wait()
}
// SyncMapConcurrentWrite 并发写
func (sm *SyncMap) SyncMapConcurrentWrite() {
	wg.Add(2)
	go sm.SyncMapWrite()
	go sm.SyncMapWrite()
	wg.Wait()
}
// SyncMapConcurrentWrite 并发写
func (sm *SyncMap) SyncMapReadAndWrite() {
	wg.Add(2)
	go sm.SyncMapRead()
	go sm.SyncMapWrite()
	wg.Wait()
}

// 使用map，实现map
func main() {

	sm := SyncMap{}
	sm.SyncMapConcurrentRead()

	//mmp := MutexMap{m: make(map[string]string)}
	//mmp.ReadAndWrite()
	//mmp.ConcurrentRead()
	//mmp.ConcurrentReadLock()
	//mmp.ConcurrentWrite()

	//fmt.Println(mp["test"])
	////delete(mp,"test")
	//fmt.Println(mp["test"])
	//for k, v := range mp {
	//	fmt.Println("k:", k)
	//	fmt.Println("v:", v)
	//}
	//
	//str, exist := mp["test"]
	//if exist {
	//	fmt.Println(str)
	//}

	//if mp["test"] == mp["noExist"] { //key "noExist"不存在，读取时输出默认值
	//	mp["test"] = "fault2"
	////}
	//fmt.Println(mp["test"])

	//mp2 := new(map[string]string)
	//*mp2 = map[string]string{}
	//(*mp2)["test2"] = "default"
	//(*mp2)["test"] = "default"
	//fmt.Println((*mp2)["test"])

	//value := new(int)
	//fmt.Print(*value)
}
