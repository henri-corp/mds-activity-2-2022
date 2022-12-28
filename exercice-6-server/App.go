package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var currentId = 1

var hostname = ""

func main() {
	hostAndPort := getEnv("HOST_AND_PORT", "0.0.0.0:9000")
	hostname, _ = os.Hostname()

	l, err := net.Listen("tcp", hostAndPort)
	if err != nil {
		return
	}

	defer l.Close()

	fmt.Println("starting server on " + hostAndPort)

	for {

		conn, err := l.Accept()
		if err != nil {
			return
		}

		go handleUserConnection(conn, currentId)
		currentId++
	}

}

func handleUserConnection(conn net.Conn, userId int) {

	defer log.Printf("#%v left", userId)
	defer conn.Close()
	log.Printf("#%v has joined", userId)
	for {
		userInput, err := bufio.NewReader(conn).ReadString('\n')
		userInput = strings.TrimSuffix(userInput, "\n")
		if err != nil {
			return
		}
		log.Printf("#%v   [IN]    %v", userId, userInput)
		log.Printf("#%v   [OUT]   PONG: %v", userId, hostname)

		conn.Write([]byte("PONG: " + hostname))
	}
}

func getEnv(key string, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return val
}
