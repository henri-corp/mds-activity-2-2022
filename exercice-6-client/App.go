package main

import (
	_ "embed"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var myName = ""

func main() {
	myName, _ = os.Hostname()
	server := getEnv("SERVER", "0.0.0.0:9000")

	tcpAddr, err := net.ResolveTCPAddr("tcp", server)
	if err != nil {
		log.Fatal("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal("Dial failed:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	log.Printf("[Server %v started]", myName)

	for {

		log.Printf("Sending: PING %v", myName)
		_, err = conn.Write([]byte(fmt.Sprintf("PING: %v\n", myName)))
		if err != nil {
			log.Fatal("Write to server failed:", err.Error())
			os.Exit(1)
		}

		reply := make([]byte, 1024)
		_, err = conn.Read(reply)
		if err != nil {
			log.Fatal("Write to server failed:", err.Error())
			os.Exit(1)
		}

		log.Printf("Reply received: %v", string(reply))
		time.Sleep(5 * time.Second)
	}
}
func getEnv(key string, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return val
}
