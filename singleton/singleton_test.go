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
		if s := GetInstanceEager(); s == nil {
			b.Fatalf("singleton_test::BenchmarkEager - No instance returned, %s", s)
		}
	}
}

func BenchmarkLazy(b *testing.B){
	for i := 0; i < b.N; i++ {
		if s := GetInstanceLazy(); s == nil {
			b.Fatalf("singleton_test::BenchmarkLazy - No instance returned, %s", s)
		}
	}
}

func BenchmarkThreadSafe(b *testing.B){
	for i := 0; i < b.N; i++ {
		if s := GetInstanceThreadSafe(); s == nil {
			b.Fatalf("singleton_test::BenchmarkThreadSafe - No instance returned, %s", s)
		}
	}
}