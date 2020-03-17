package main

import (
	"encoding/json"	
	"log"
	"net/http"
)

type request struct {
	Headers map[string][]string `json:"headers"`
}

func datafy(r *http.Request) *request {
	headers := make(map[string][]string)
	for name, value := range r.Header {
		headers[name] = value
	}
	return &request{
		Headers: headers,
	}
}

func echoRequest(w http.ResponseWriter, r *http.Request) {
	request := datafy(r)

	encoder := json.NewEncoder(w)
	err := encoder.Encode(&request)
	if err != nil {
		log.Printf("error echoing request: %v", err)
	}
}

func main() {
	http.HandleFunc("/", echoRequest)

	log.Print("started server at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
