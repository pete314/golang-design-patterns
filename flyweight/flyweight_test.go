// Author: Peter Nagy <https://peternagy.ie>
// Since: Dec, 2017
// Package: flyweight
// Description: flyweight patter benchmark tests

package flyweight

import "testing"

var (
	colors = []string{"red", "blue", "green", "yellow", "cyan", "white"}
)

func BenchmarkFlyweight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if c := ColorFactory(colors[i%6]); c == nil {
			b.Fatal("flyweight_test::BenchmarkFlyweight - No instance returned")
		}
	}
}

func BenchmarkNewInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if c := NewColor(colors[i%6]); c == nil {
			b.Fatal("flyweight_test::NewInstance - No instance returned")
		}
	}
}
