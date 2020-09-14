package main

import "testing"

func BenchmarkPrintArgs(b *testing.B) {
	testStrSlice := []string{
		"hello",
		"world",
		"你好",
		"世界",
	}
	for i := 0; i < b.N; i++ {
		PrintArgs(testStrSlice)
	}
}
