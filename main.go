package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Could not find data file. Are you sure data file exists?")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	
	matches := regexp.MustCompile(`^\d+.*$`)
	nums := []float64{}
	for scanner.Scan() {
		line := scanner.Text()

		match := matches.FindString(line)
		if match != "" {
			x, err := strconv.ParseFloat(match, 64)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}

			nums = append(nums, x)
		}

	}

	if len(nums) == 0 {
		fmt.Println("Average: 0")
		fmt.Println("Median: 0")
		fmt.Println("Variance: 0")
		fmt.Println("Standard Deviation: 0")

		os.Exit(0)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	avg := averageNum(nums)
	med := Median(nums)
	vari := Variance(nums, avg)
	std := math.Sqrt(vari)
	fmt.Println("Average:", math.Round(avg))
	fmt.Println("Median:", med)
	fmt.Println("Variance:", math.Round(vari))
	fmt.Println("Standard Deviation:", math.Round(std))

}
