package main

import (
	"log"
	"net/http"
)

func main()  {

	oteApi := NewOteApi()

	//setup handler routes
	http.Handle("/ote/electricity", oteApi.GetElectricityOteData())
	http.Handle("/ote/electricity/evaluate", oteApi.EvaluateElectricityValue())

	log.Println("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}