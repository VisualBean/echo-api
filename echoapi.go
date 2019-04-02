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
	port = flag.String("port", "8000", "The port to run the service on. (default: 8000)")
)

type input struct {
	Method string
	Body   interface{}
}

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		outputRequestdata(r)
	})

	go http.ListenAndServe(":"+*port, nil)

	log.Println("Echo-API started.", "Listening on", *port, "press <ENTER> to exit")
	fmt.Scanln()
}

func outputRequestdata(r *http.Request) {
	bodyJSON, _ := ioutil.ReadAll(r.Body)

	requestData := input{Method: r.Method,
		Body: new(interface{}),
	}

	if err := json.Unmarshal(bodyJSON, requestData.Body); err != nil {
		fmt.Println(err)
	}

	output, err := json.MarshalIndent(requestData, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "at", r.URL)
	fmt.Print(string(output))
	fmt.Println("\n")
}
