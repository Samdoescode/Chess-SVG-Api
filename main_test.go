package main

import (
	"testing"
)

var num = 1000

func BenchmarkMain(b *testing.B){
	for i :=0; i < b.N; i++ {
		main()
	}
}