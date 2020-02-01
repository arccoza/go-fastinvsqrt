package fis

import (
	// "fmt"
	"math"
	"testing"
)

var result32 float32
var result64 float64

// func init() {
// 	fmt.Println(invSqrt64(1.6))
// 	fmt.Println(FastInvSqrt32(1.6))
// 	fmt.Println(FastInvSqrt64(1.6))

// 	fmt.Println(invSqrt64(-1.6))
// 	fmt.Println(FastInvSqrt32(-1.6))
// 	fmt.Println(FastInvSqrt64(-1.6))
// }

func TestFastInvSqrt32(t *testing.T) {
	a, b := float64(FastInvSqrt32(1.6)), invSqrt64(1.6)
	got := float32(math.Abs(a - b))
	epsilon := 0.002
	error := float32(math.Max(1.0, math.Max(a, b)) * epsilon)

	if got > error {
		t.Errorf("FastInvSqrt32(1.6) error is %v; want %v or less", got, epsilon)
	}
}

func TestFastInvSqrt64(t *testing.T) {
	a, b := FastInvSqrt64(1.6), invSqrt64(1.6)
	got := math.Abs(a - b)
	epsilon := 0.002
	error := math.Max(1.0, math.Max(a, b)) * epsilon

	if got > error {
		t.Errorf("FastInvSqrt64(1.6) error is %v; want %v or less", got, epsilon)
	}
}

func TestFastInvSqrt32ForNegative(t *testing.T) {
	got := FastInvSqrt32(-1)
	if !math.IsNaN(float64(got)) {
		t.Errorf("FastInvSqrt32(-1) is %v; want NaN", got)
	}
}

func TestFastInvSqrt64ForNegative(t *testing.T) {
	got := FastInvSqrt64(-1)
	if !math.IsNaN(got) {
		t.Errorf("FastInvSqrt64(-1) is %v; want NaN", got)
	}
}

func benchmarkInvSqrt64(val float64, b *testing.B) {
	var r float64
	for i := 0; i < b.N; i++ {
		r = invSqrt64(val)
	}
	result64 = r
}

func BenchmarkInvSqrt64_1(b *testing.B) { benchmarkInvSqrt64(1.6, b) }
func BenchmarkInvSqrt64_2(b *testing.B) { benchmarkInvSqrt64(math.Pi, b) }
func BenchmarkInvSqrt64_3(b *testing.B) { benchmarkInvSqrt64(math.E, b) }
func BenchmarkInvSqrt64_4(b *testing.B) { benchmarkInvSqrt64(5.3, b) }
func BenchmarkInvSqrt64_5(b *testing.B) { benchmarkInvSqrt64(4, b) }

func benchmarkFastInvSqrt32(val float32, b *testing.B) {
	var r float32
	for i := 0; i < b.N; i++ {
		r = FastInvSqrt32(val)
	}
	result32 = r
}

func BenchmarkFastInvSqrt32_1(b *testing.B) { benchmarkFastInvSqrt32(1.6, b) }
func BenchmarkFastInvSqrt32_2(b *testing.B) { benchmarkFastInvSqrt32(math.Pi, b) }
func BenchmarkFastInvSqrt32_3(b *testing.B) { benchmarkFastInvSqrt32(math.E, b) }
func BenchmarkFastInvSqrt32_4(b *testing.B) { benchmarkFastInvSqrt32(5.3, b) }
func BenchmarkFastInvSqrt32_5(b *testing.B) { benchmarkFastInvSqrt32(4, b) }

func benchmarkFastInvSqrt64(val float64, b *testing.B) {
	var r float64
	for i := 0; i < b.N; i++ {
		r = FastInvSqrt64(val)
	}
	result64 = r
}

func BenchmarkFastInvSqrt64_1(b *testing.B) { benchmarkFastInvSqrt64(1.6, b) }
func BenchmarkFastInvSqrt64_2(b *testing.B) { benchmarkFastInvSqrt64(math.Pi, b) }
func BenchmarkFastInvSqrt64_3(b *testing.B) { benchmarkFastInvSqrt64(math.E, b) }
func BenchmarkFastInvSqrt64_4(b *testing.B) { benchmarkFastInvSqrt64(5.3, b) }
func BenchmarkFastInvSqrt64_5(b *testing.B) { benchmarkFastInvSqrt64(4, b) }
