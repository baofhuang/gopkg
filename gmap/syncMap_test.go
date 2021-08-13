package main

import "testing"

func BenchmarkSyncMap_SyncMapConcurrentRead(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm := SyncMap{}
		sm.SyncMapConcurrentRead()
	}

}
func BenchmarkSyncMap_SyncMapConcurrentWrite(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm := SyncMap{}
		sm.SyncMapConcurrentWrite()
	}
}
func BenchmarkSyncMap_SyncMapReadAndWrite(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm := SyncMap{}
		sm.SyncMapReadAndWrite()
	}
}
