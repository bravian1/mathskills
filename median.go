package main

import "sort"

func Median(n []float64) int {
	sort.Float64s(n)

	if len(n) == 0 {
		return 0
	}
	l := len(n)
	if l%2 == 0 {
		return int(n[l/2-1]+n[l/2]) / 2
	}
	return int(n[l/2])
}
