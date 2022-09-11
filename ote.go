package main

import "time"

//ElectricityHourData represent electricity data on hourly base
type ElectricityHourData struct {
	Time string `json:"time"`
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	Avg float64 `json:"avg"`
	Value float64 `json:"value"`
}

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
type ElectricityDataService interface {
	GetElectricityHourData() (*ElectricityHourData, error)
}