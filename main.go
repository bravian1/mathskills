package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	nums := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		matches := regexp.MustCompile(`^\d+.*$`)
		match := matches.FindAllString(line, -1)
		fmt.Println(match)
		for _, char := range match {
			nums = append(nums, char)
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	num := []float64{}
	for _, char := range nums {
		x, err := strconv.ParseFloat(char, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		num = append(num, x)

	}

	//fmt.Println("Numbers:", nums)
	fmt.Println("Average:", averageNum(num))
	fmt.Println("Median:", Median(num))

}
