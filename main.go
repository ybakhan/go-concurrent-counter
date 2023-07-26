package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var counter int = 0

func get(writer http.ResponseWriter, req *http.Request) {
	log.Printf("GET counter request: %v", counter)
	fmt.Fprintf(writer, "Counter is at: %d\n", counter)
}

func set(writer http.ResponseWriter, req *http.Request) {
	log.Printf("SET counter request: %v", req.RequestURI)
	value := req.URL.Query().Get("value")
	intval, err := strconv.Atoi(value)

	if err != nil {
		log.Println("SET handler: non-integer parameter value.")
	}

	counter = intval
	log.Printf("counter set to: %v", counter)
	fmt.Fprintf(writer, "Counter set to: %d\n", counter)
}

func inc(_ http.ResponseWriter, _ *http.Request) {
	counter = counter + 1
	log.Printf("counter incremented to: %v", counter)
}

func main() {
	http.HandleFunc("/counter", get)
	http.HandleFunc("/counter/set", set)
	http.HandleFunc("/increment", inc)

	port := 9095
	if len(os.Args) > 1 {
		port, _ = strconv.Atoi(os.Args[1])
	}
	log.Printf("Listening on port %d\n", port)
	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(port), nil))
}
