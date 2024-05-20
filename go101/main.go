package main

import u "unsafe"

type A struct {x int}
type B struct { x *int}

func main(){
  if u.Sizeof(A{}) == u.Sizeof(B{}) {
    x := append(make([]A, 33), A{})
    y := append(make([]B, 33), B{})
    println(cap(x) == cap(y))
  }
}
