package main

import "testing"

func TestAdd(t *testing.T) {
	t.Logf("add")
}

func TestMain(m *testing.M) {
	m.Run()
}
