package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var counter int = 0

func get(w http.ResponseWriter, req *http.Request) {
	log.Printf("GET received: %v", counter)
	fmt.Fprintf(w, "got: %d\n", counter)
}

func set(w http.ResponseWriter, req *http.Request) {
	log.Printf("SET counter to: %v", req)
	val := req.URL.Query().Get("val")
	intval, err := strconv.Atoi(val)

	if err != nil {
		panic("unhandled error in API SET handler")
	}

	counter = intval
	log.Printf("set to: %v", counter)

}

func inc(_ http.ResponseWriter, _ *http.Request) {
	counter = counter + 1
	log.Printf("incremented to: %v", counter)
}

func main() {
	http.HandleFunc("/get", get)
	http.HandleFunc("/set", set)
	http.HandleFunc("/increment", inc)

	portnum := 9095
	if len(os.Args) > 1 {
		portnum, _ = strconv.Atoi(os.Args[1])
	}
	log.Printf("Listening on port %d\n", portnum)
	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(portnum), nil))
}