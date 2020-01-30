package fis

import (
  "math"
)

const (
  magic32 = 0x5F375A86
  magic64 = 0x5FE6EB50C7B537A9
)

func invSqrt64(n float64) float64 {
  return 1 / math.Sqrt(n)
}

func FastInvSqrt32(n float32) float32 {
  n2 := n * 0.5
  th := float32(1.5)
  b := math.Float32bits(n)
  b = magic32 - (b >> 1)
  f := math.Float32frombits(b)
  f *= th - (n2 * f * f)
  return f
}

func FastInvSqrt64(n float64) float64 {
  n2 := n * 0.5
  th := float64(1.5)
  b := math.Float64bits(n)
  b = magic64 - (b >> 1)
  f := math.Float64frombits(b)
  f *= th - (n2 * f * f)
  return f
}