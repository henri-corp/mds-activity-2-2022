package main

import (
	_ "embed"
	"fmt"
	"net/http"
)

const HOST_AND_PORT = "0.0.0.0:9000"

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("starting to listen on " + HOST_AND_PORT)
	http.ListenAndServe(HOST_AND_PORT, nil)

}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ex4: It works!\n")

}
