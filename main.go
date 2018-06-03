package main

import (
	"time"
	"net/http"
	"log"
	"fmt"
)

func main() {
	url := "http://googleleandro.com"

	if err := waitForServer(url); err != nil {
		log.Fatalf("Site is down: %v", err)
	}
	fmt.Println("execution successful")
}

func waitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		fmt.Println(tries)
		_, err := http.Head(url)
		if err == nil {
			return nil
		}

		log.Printf("server not responding (%s); retrying...", err)

		// increment 2x tries
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
