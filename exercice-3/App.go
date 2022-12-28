package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
)

const filename = "/data/motd.txt"

func main() {

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("ERROR", "impossible de trouver le fichier", filename, "dans le conteneur")
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}
	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}

}
