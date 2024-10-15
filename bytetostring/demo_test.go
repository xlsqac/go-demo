package main

import (
	"runtime"
	"testing"
	"unsafe"
)

var by = []byte("test")

func bytesToString(bs []byte) string {
	return string(bs)
}

func bytesToString1(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

func bytesToString2(bs []byte) string

func bytesToString3(bs []byte) string {
	return unsafe.String(unsafe.SliceData(bs), len(bs))
}

func BenchmarkBase(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = bytesToString(by)
	}
	runtime.KeepAlive(s)
}

func BenchmarkAssembly(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = bytesToString2(by)
	}
	runtime.KeepAlive(s)
}

func BenchmarkUnsafe(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = bytesToString3(by)
	}
	runtime.KeepAlive(s)
}

func BenchmarkUnsafeP(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = bytesToString1(by)
	}
	runtime.KeepAlive(s)
}

