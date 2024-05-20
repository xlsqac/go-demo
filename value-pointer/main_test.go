package main

import "testing"

func BenchmarkStructReturnValue(b *testing.B) {
	b.ReportAllocs()
	testS := testStruct{field: "testStruct"}

	for i := 0; i < b.N; i++ {
		value(testS)
	}
}

func BenchmarkStructReturnPointer(b *testing.B) {
	b.ReportAllocs()
	testS := testStruct{field: "testStruct"}

	for i := 0; i < b.N; i++ {
		pointer(&testS)
	}
}
