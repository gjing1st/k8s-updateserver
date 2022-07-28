package utils

import (
	"testing"
)

var ifc interface{}

func BenchmarkInt(b *testing.B) {
	ifc = "123577"
	for i := 0; i < b.N; i++ {
		Int(ifc)
	}
}

func BenchmarkString(b *testing.B) {
	ifc = "sdsd"
	for i := 0; i < b.N; i++ {
		String(ifc)
	}
}