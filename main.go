package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main()  {

	http.HandleFunc("/ote/electricity", func(w http.ResponseWriter, r *http.Request){

		w.Header().Set("Content-Type", "application/json")
		newElectricityDataService := NewHtmlElectricityDataService()
		res, err := newElectricityDataService.GetElectricityHourData()

		if err != nil {
			http.Error(w, fmt.Sprintf("error get ote electricity data: %s", err.Error()), http.StatusInternalServerError)
		}

		err = json.NewEncoder(w).Encode(res)

		if err != nil {
			http.Error(w, fmt.Sprintf("error encode json ote electricity data: %s", err.Error()), http.StatusInternalServerError)
		}
	})

	log.Println("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}