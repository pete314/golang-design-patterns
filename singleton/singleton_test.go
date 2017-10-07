// Author: Peter Nagy <https://peternagy.ie>
// Since: Oct, 2017
// Package: singleton
// Description: benchmark test for eager singleton
package singleton

import (
	"testing"
)

func BenchmarkEager(b *testing.B){
	for i := 0; i < b.N; i++ {
		if e := GetInstanceEager(); e == nil {
			b.Fatalf("No instance returned", e)
		}
	}
}

func BenchmarkLazy(b *testing.B){
	for i := 0; i < b.N; i++ {
		if e := GetInstanceLazy(); e == nil {
			b.Fatalf("No instance returned", e)
		}
	}
}