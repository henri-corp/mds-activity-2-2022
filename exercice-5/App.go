package main

import (
	_ "embed"
	"fmt"
	"net/http"
	"os"
)

func main() {
	hostAndPort := getEnv("HOST_AND_PORT", "127.0.0.1:9000")
	http.HandleFunc("/", handler)
	fmt.Println("starting to listen container on " + hostAndPort)
	http.ListenAndServe(hostAndPort, nil)

}

func getEnv(key string, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return val
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ex5: It works!\n")

}
