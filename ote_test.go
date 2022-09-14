package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MockServerWithHandler(handler http.Handler) (*httptest.Server, error) {
	// create a listener with the desired port.
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		return nil, err
	}

	w := httptest.NewUnstartedServer(handler)
	w.Listener.Close()
	w.Listener = l
	w.Start()

	return w, nil
}

func TestGetOteData(t *testing.T) {
	oteApi := NewOteApi()
	currentVal, err := NewHtmlElectricityDataService().GetElectricityHourData()
	if err != nil {
		t.Error(err)
	}

	httpServer, err := MockServerWithHandler(oteApi.GetElectricityOteData())
	if err != nil {
		t.Error(err)
	}
	defer httpServer.Close()

	res, err := httpServer.Client().Get("http://localhost:8080/ote/electricity")

	if err != nil {
		t.Error(err)
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Error(err)
	}

	var newRes *ElectricityHourData
	err = json.Unmarshal(data, &newRes)

	if err != nil {
		t.Error(err)
	}

	if newRes.Value != currentVal.Value {
		t.Errorf("the electricity hour data is not equal")
	}
}

//TestEvaluateElectricityValue test evaluation
func TestEvaluateElectricityValue(t *testing.T) {

	oteApi := NewOteApi()
	randomVal := &ElectricityValue{Value: rand.Float64()}
	currentVal, err := NewHtmlElectricityDataService().GetElectricityHourData()
	if err != nil {
		t.Error(err)
	}

	expected := EvaluateElectricityData(randomVal, currentVal)

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(randomVal)
	if err != nil {
		t.Error(err)
	}

	httpServer, err := MockServerWithHandler(oteApi.EvaluateElectricityValue())
	if err != nil {
		t.Error(err)
	}
	defer httpServer.Close()

	res, err := httpServer.Client().Post("http://localhost:8080/ote/electricity/evaluate", "application/json", &buf)

	if err != nil {
		t.Error(err)
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Error(err)
	}

	var newRes ElectricityValueResponse
	err = json.Unmarshal(data, &newRes)

	if err != nil {
		t.Error(err)
	}

	if newRes.Value != expected {
		t.Errorf("the evaluation was success. Expected %v and get %v", expected, newRes.Value)
	}
}
