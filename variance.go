package main

import "math"

func Variance(n []float64, avg float64) float64 {
	sum := float64(0)

	for i := 0; i < len(n); i++ {
		sum += math.Pow((n[i] - avg), 2)
	}
	va := sum / float64(len(n))

	return va

}
