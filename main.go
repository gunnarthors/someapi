package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := 8080

	http.HandleFunc("/ping", Ping)
	http.HandleFunc("/somedata", GetData)

	addr := ":" + strconv.Itoa(port)
	log.Printf("Listening on %q", addr)
	log.Fatal(http.ListenAndServe(addr, nil))

}

// Ping is a function
func Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	_, _ = w.Write([]byte("yes, this is api that has a working ping..."))
}

// GetData is a function
func GetData(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	_, _ = w.Write([]byte("Here is a environment variable:" + os.Getenv("ENVVAR")))
}
