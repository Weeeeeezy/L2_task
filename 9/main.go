package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get(os.Args[1])
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("Invalid site")
	}

	recordFile("index.html", resp)
}

func recordFile(file string, m *http.Response) {
	outFile, err := os.Create(file)
	defer outFile.Close()
	if err != nil {
		log.Print("Cannot create file")
	}

	body, err := ioutil.ReadAll(m.Body)
	if err != nil {
		log.Println("Cannot read site body")
	}
	outFile.WriteString(string(body))
}
