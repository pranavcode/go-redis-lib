package main

import (
	"fmt"

	"urlbase"
	"helper/urls"
)

func main() {
	// Initialize urlbase library
	err := urlbase.Initialize("192.168.1.10:6379")
	if !err {
		// Finalize urlbase
		defer urlbase.Finalize()

		fmt.Println("Welcome to demonstration for go-lib")
		
		// Generate URLs
		fmt.Println("\nGenerating 10 random URLs for the demo...")
		stored_urls := urls.Generate(20, 10)
		fmt.Println(stored_urls)

		// Store them using urlbase library
		fmt.Println("\nStoring the URLs into urlbase (Redis DB client code)...")
		for i := 0; i < len(stored_urls); i++ {
			check := urlbase.Store(i, stored_urls[i])
			if !check {
				fmt.Println("Failed to store", stored_urls[i])				
			}
			fmt.Println("Stored", stored_urls[i])
		}
		fmt.Println("Stored",len(stored_urls),"URLs!")

		// Retrieve stored URLs using urlbase library
		retrieved_urls := make([]string, len(stored_urls))
		fmt.Println("\nRetrieving stored URLs from Redis DB...")
		for i := 0; i < len(retrieved_urls); i++ {
			value, err := urlbase.Fetch(i)
			if !err {				
				retrieved_urls[i] = value
				fmt.Println("Retrieved", retrieved_urls[i],
					"Matched", (retrieved_urls[i] == stored_urls[i]))
			}
		}
		fmt.Println("Retrieved",len(retrieved_urls),"URLs!")
	}
}
