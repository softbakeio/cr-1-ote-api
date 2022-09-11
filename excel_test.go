package main

import "testing"

func TestExcelElectricityDataServiceNoEmpty(t *testing.T) {
	newExcelElectricityService := NewExcelElectricityDataService()
	val, err := newExcelElectricityService.GetElectricityHourData()

	if err != nil {
		t.Error(err)
	}

	if val == nil {
		t.Errorf("electricity data si nil")
	}
}

