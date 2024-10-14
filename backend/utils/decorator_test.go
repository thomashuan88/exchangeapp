package utils

import "testing"

type ExecuteFn func(string)

func Execute(fn ExecuteFn) {
	fn("hello")
}

func TestDecoratorxx1(t *testing.T) {
	Execute(func(s string) {
		t.Log(s)
	})
}
