package main

import "math"
/*
function that calculates the variance of a slice of foats and returns the variance by 
the formula variance = ((value-average)^2)/number of values
*/
func Variance(n []float64, avg float64) float64 {
	sum := float64(0)

	for i := 0; i < len(n); i++ {
		sum += math.Pow((n[i] - avg), 2)
	}
	vari := sum / float64(len(n))

	return vari

}
