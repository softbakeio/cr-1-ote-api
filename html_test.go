package main

import (
	"testing"
)

//TestHtmlElectricityHourDataNotEmpty testing if html parsing electricity data
// not thrown error or nil
func TestHtmlElectricityHourDataNotEmpty(t *testing.T) {
	newHtmlElectricityService := NewHtmlElectricityDataService()
	val, err := newHtmlElectricityService.GetElectricityHourData()

	if err != nil {
		t.Error(err)
	}

	if val == nil {
		t.Errorf("electricity data si nil")
	}
}