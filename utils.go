package main

import (
	"strings"
	"time"
)

func NormalizeAmerican(old string) string {
	return strings.TrimSpace(strings.Replace(old, ",", ".", -1))
}

func GetCurrentTimestamp() string {
	return time.Now().Format(time.RFC850)
}

// MinMax find min and max float number in array
func MinMax(array []float64) (float64, float64) {
	var max float64 = array[0]
	var min float64 = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

//Avg calculate avg of float64 numbers
func Avg(array []float64) float64 {
	// declaring a variable
	// to store the sum
	var sum float64 = 0
	// traversing through the
	// array using for loop
	for i := 0; i < len(array); i++ {

		// adding the values of
		// array to the variable sum
		sum += array[i]
	}

	// declaring a variable
	// avg to find the average
	return  sum / float64(len(array))
}