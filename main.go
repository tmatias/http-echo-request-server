package main

import (
	"encoding/json"	
	"io/ioutil"
	"log"
	"net/http"
)

type request struct {
	Method		string `json:"method"`
	URL			string `json:"url"`
	Headers 	map[string][]string `json:"headers"`
	Body		string `json:"body"`
	RemoteAddr 	string `json:"remote_address"`
}

func datafy(r *http.Request) *request {
	headers := make(map[string][]string)
	for name, value := range r.Header {
		headers[name] = value
	}
	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading request body: %v\n", err)
		bs = []byte("<error reading body>")
	}
	return &request{
		Method: r.Method,
		URL: r.URL.String(),
		Headers: headers,
		Body: string(bs),
		RemoteAddr: r.RemoteAddr,
	}
}

func echoRequest(w http.ResponseWriter, r *http.Request) {
	request := datafy(r)

	encoder := json.NewEncoder(w)
	err := encoder.Encode(&request)
	if err != nil {
		log.Printf("error echoing request: %v\n", err)
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
