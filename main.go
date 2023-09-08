package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var counter = Counter{}

type Counter struct {
	mu    sync.Mutex
	value int
}

func get(writer http.ResponseWriter, req *http.Request) {
	counter.mu.Lock()
	defer counter.mu.Unlock()
	log.Printf("GET counter request: %v", counter.value)
	fmt.Fprintf(writer, "Counter is at: %d\n", counter.value)
}

func set(writer http.ResponseWriter, req *http.Request) {
	counter.mu.Lock()
	defer counter.mu.Unlock()
	log.Printf("SET counter request: %v", req.RequestURI)
	value := req.URL.Query().Get("value")
	intval, err := strconv.Atoi(value)

	if err != nil {
		log.Println("SET handler: non-integer parameter value.")
	}

	counter.value = intval
	log.Printf("counter set to: %v", counter.value)
	fmt.Fprintf(writer, "Counter set to: %d\n", counter.value)
}

func inc(_ http.ResponseWriter, _ *http.Request) {
	counter.mu.Lock()
	defer counter.mu.Unlock()
	counter.value++
	log.Printf("counter incremented to: %v", counter.value)
}

func dec(_ http.ResponseWriter, _ *http.Request) {
	counter.mu.Lock()
	defer counter.mu.Unlock()
	counter.value--
	log.Printf("counter decremented to: %v", counter.value)
}

func main() {
	http.HandleFunc("/counter", get)
	http.HandleFunc("/counter/set", set)
	http.HandleFunc("/increment", inc)
	http.HandleFunc("/decrement", dec)

	port := 9095
	if len(os.Args) > 1 {
		port, _ = strconv.Atoi(os.Args[1])
	}
	log.Printf("Listening on port %d\n", port)
	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(port), nil))
}
