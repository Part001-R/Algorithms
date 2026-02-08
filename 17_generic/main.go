//go:build !deprecated
// +build !deprecated

package main

import "fmt"

type mytype interface {
	~int | ~string
}

type Foo struct{}

func testGeneric[T mytype, V int](key T, val V) map[T]V {
	m := make(map[T]V, 0)
	m[key] = val
	return m
}

// aaa is deprecated, use bbb instead
//
// Deprecated: Use bbb instead
func aaa() {
	fmt.Println("Hello")
}

func main() {
	m := testGeneric("сто", 100)
	fmt.Println(m)

	aaa()
}
