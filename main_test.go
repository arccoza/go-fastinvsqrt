package main

import "testing"

const val = 1.6
var result32 float32
var result64 float64

func BenchmarkInvSqrt64(b *testing.B) {
	var r float64
    for i := 0; i < b.N; i++ {
        r = InvSqrt64(val)
    }
	result64 = r
}

func BenchmarkFastInvSqrt32(b *testing.B) {
	var r float32
    for i := 0; i < b.N; i++ {
        r = FastInvSqrt32(val)
    }
	result32 = r
}

func BenchmarkFastInvSqrt64(b *testing.B) {
	var r float64
    for i := 0; i < b.N; i++ {
        r = FastInvSqrt64(val)
    }
	result64 = r
}