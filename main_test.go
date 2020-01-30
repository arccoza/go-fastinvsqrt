package main

import (
	"testing"
	"math"
)

const val = 1.6
var result32 float32
var result64 float64

func benchmarkInvSqrt64(val float64, b *testing.B) {
	var r float64
    for i := 0; i < b.N; i++ {
        r = InvSqrt64(val)
    }
	result64 = r
}

func BenchmarkInvSqrt64_1(b *testing.B)  { benchmarkInvSqrt64(1.6, b) }
func BenchmarkInvSqrt64_2(b *testing.B)  { benchmarkInvSqrt64(math.Pi, b) }
func BenchmarkInvSqrt64_3(b *testing.B)  { benchmarkInvSqrt64(math.E, b) }
func BenchmarkInvSqrt64_4(b *testing.B)  { benchmarkInvSqrt64(5.3, b) }
func BenchmarkInvSqrt64_5(b *testing.B)  { benchmarkInvSqrt64(4, b) }


func benchmarkFastInvSqrt32(val float32, b *testing.B) {
	var r float32
    for i := 0; i < b.N; i++ {
        r = FastInvSqrt32(val)
    }
	result32 = r
}

func BenchmarkFastInvSqrt32_1(b *testing.B)  { benchmarkFastInvSqrt32(1.6, b) }
func BenchmarkFastInvSqrt32_2(b *testing.B)  { benchmarkFastInvSqrt32(math.Pi, b) }
func BenchmarkFastInvSqrt32_3(b *testing.B)  { benchmarkFastInvSqrt32(math.E, b) }
func BenchmarkFastInvSqrt32_4(b *testing.B)  { benchmarkFastInvSqrt32(5.3, b) }
func BenchmarkFastInvSqrt32_5(b *testing.B)  { benchmarkFastInvSqrt32(4, b) }


func benchmarkFastInvSqrt64(val float64, b *testing.B) {
	var r float64
    for i := 0; i < b.N; i++ {
        r = FastInvSqrt64(val)
    }
	result64 = r
}

func BenchmarkFastInvSqrt64_1(b *testing.B)  { benchmarkFastInvSqrt64(1.6, b) }
func BenchmarkFastInvSqrt64_2(b *testing.B)  { benchmarkFastInvSqrt64(math.Pi, b) }
func BenchmarkFastInvSqrt64_3(b *testing.B)  { benchmarkFastInvSqrt64(math.E, b) }
func BenchmarkFastInvSqrt64_4(b *testing.B)  { benchmarkFastInvSqrt64(5.3, b) }
func BenchmarkFastInvSqrt64_5(b *testing.B)  { benchmarkFastInvSqrt64(4, b) }