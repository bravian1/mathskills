package main
/*
function that calculates the average of a slice of float64
by dividing the sum of all the numbers by the number of values
*/
func averageNum(n []float64) float64 {
	sum := float64(0)
	for i := 0; i < len(n); i++ {
		sum = sum + n[i]
	}
	avg := (float64(sum)) / (float64(len(n)))
	
	return (avg)
}
