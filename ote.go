package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (

)

//OteApi struct of OTE API
type OteApi struct {
	newElectricityDataService ElectricityDataService
}

//NewOteApi create new instance of OteApi
func NewOteApi() *OteApi{
	newInstance := new(OteApi)
	newInstance.newElectricityDataService = NewHtmlElectricityDataService()
	return newInstance
}

//ElectricityHourData represent electricity data on hourly base
type ElectricityHourData struct {
	Time string `json:"time"`
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	Avg float64 `json:"avg"`
	Value float64 `json:"value"`
}

type ElectricityValue struct {
	Value float64 `json:"value"`
}

type ElectricityValueResponse struct {
	Value bool `json:"value"`
}


//SetupApiHeaders setuo api headers for each endpoint
func (oa *OteApi) SetupApiHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

// EvaluateElectricityValue evaluate user electricity value with the current OTE
// market values
func (oa *OteApi) EvaluateElectricityValue() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var electricityValue *ElectricityValue
		oa.SetupApiHeaders(w)

		if r.Method != "POST" {
			http.Error(w, "only POST is supported", http.StatusBadRequest)
			return
		}

		err := json.NewDecoder(r.Body).Decode(&electricityValue)

		if err != nil {
			http.Error(w, fmt.Sprintf("error read electricity value: %s", err.Error()), http.StatusBadRequest)
			return
		}

		result, err := oa.newElectricityDataService.GetElectricityHourData()

		if err != nil {
			http.Error(w, fmt.Sprintf("error get ote electricity data: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		eval := EvaluateElectricityData(electricityValue, result)

		err = json.NewEncoder(w).Encode(&ElectricityValueResponse{Value: eval})

		if err != nil {
			http.Error(w, fmt.Sprintf("error encode json ote electricity value response: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	})
}

// GetElectricityOteData get current electricity OTE data
func (oa *OteApi) GetElectricityOteData() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			http.Error(w, "only GET is supported", http.StatusBadRequest)
			return
		}

		oa.SetupApiHeaders(w)
		res, err := oa.newElectricityDataService.GetElectricityHourData()

		if err != nil {
			http.Error(w, fmt.Sprintf("error get ote electricity data: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(res)

		if err != nil {
			http.Error(w, fmt.Sprintf("error encode json ote electricity data: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	})
}

type ElectricityDataService interface {
	GetElectricityHourData() (*ElectricityHourData, error)
}