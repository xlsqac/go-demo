package main

import (
	"testing"
	"unique"
)

func BenchmarkMake(b *testing.B) {
	s1 := "hello"
	s2 := "world"
	for i := 0; i < b.N; i++ {
		compare(s1, s2)
	}

}

func BenchmarkMake1(b *testing.B) {
	s1 := unique.Make("hello")
	s2 := unique.Make("world")
	for i := 0; i < b.N; i++ {
		compare1(s1, s2)
	}

}

func compare(s1, s2 string) bool {
	return s1 == s2
}

func compare1(s1, s2 unique.Handle[string]) bool {
	return s1 == s2
}
