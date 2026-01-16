package main

import (
	"fmt"
	"log"
	"time"

	"github.com/julian-bruyers/touchid-go"
)

func main() {
	fmt.Println("Touch ID Authentication Test")

	if !touchid.Available() {
		log.Fatal("Touch ID is not available on this system")
	}

	isAuthenticated, err := touchid.Authenticate(
		touchid.WithMsg("Verify your identity for touchid-go test"),
		touchid.WithPassword(false),
		touchid.WithTimeout(60*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}

	if isAuthenticated {
		fmt.Println("Authentication successful!")
	} else {
		fmt.Println("Authentication failed!")
	}
}
