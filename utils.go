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

//EvaluateElectricityData evaluate eletricity value with the OTE electricity hour data
// if the value is in interval of min max or is less or bigger then current hour value
func EvaluateElectricityData(electricityValue *ElectricityValue, electricityHourData *ElectricityHourData ) bool {

	// first check if the value is less then the minimal value from 24 hours interval
	if electricityValue.Value < electricityHourData.Min {
		return true
	}

	// else check if the value is bigger then the max value then return false
	if electricityValue.Value > electricityHourData.Max {
		return false
	}

	// check with the current hour value less then return true if
	// bigger then return false
	if electricityValue.Value < electricityHourData.Value {
		return true
	} else {
		return false
	}
}

// CreateElectricityHourData create new struct of ElectricityHourData
// with the data
func CreateElectricityHourData(data []float64) *ElectricityHourData {
	min, max := MinMax(data)
	avg := Avg(data)
	currentTimestamp := GetCurrentTimestamp()
	value := data[time.Now().Hour()]

	return &ElectricityHourData{
		Time: currentTimestamp,
		Avg: avg,
		Max: max,
		Min: min,
		Value: value,
	}
}