package main

import (
	"fmt"
	"log"

	"github.com/julian-bruyers/touchid-go"
)

func main() {
	fmt.Println("Touch ID Authentication Test")

	isAuthenticated, err := touchid.Authenticate("Verify your identity for touchid-go test")

	if err != nil {
		log.Fatal(err)
	}

	if isAuthenticated {
		fmt.Println("Authentication successful!")
	} else {
		fmt.Println("Authentication failed!")
	}
}
