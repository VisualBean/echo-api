package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	port           = flag.String("port", "8000", "The port to run the service on. (default: 8000)")
	staticResponse []byte
)

type input struct {
	Method string
	Body   interface{}
}

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		requestBody, _ := ioutil.ReadAll(r.Body)
		logToConsole(r.URL.Path, r.Method, requestBody)
		switch r.Method {
		case "GET":
			w.Write(staticResponse)
		case "POST":
			staticResponse = requestBody
		case "DELETE":
			staticResponse = nil
		}
	})

	go http.ListenAndServe(":"+*port, nil)

	log.Println("Echo-API started.", "Listening on", "http://localhost:"+*port, "press <ENTER> to exit")
	fmt.Scanln()
}

func logToConsole(uri string, httpMethod string, body []byte) {

	requestData := input{Method: httpMethod,
		Body: new(interface{}),
	}

	if err := json.Unmarshal(body, requestData.Body); err != nil {
		fmt.Println("Invalid json received in body.")
	}

	output, err := json.MarshalIndent(requestData, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "at", uri)
	fmt.Print(string(output))
	fmt.Println("\n")
}
