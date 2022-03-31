package main

import (
	"github.com/beevik/ntp"
	"log"
	"os"
	"time"
)

func main() {
	logger := log.New(os.Stderr, "", 0)

	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		logger.Fatal(err.Error())
		os.Exit(1)
	}

	time := time.Now().Add(response.ClockOffset)
	log.Printf("Current time: %s", time)
}
