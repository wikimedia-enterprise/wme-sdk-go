package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/wikimedia-enterprise/wme-sdk-go/pkg/auth"
)

func main() {
	client := auth.NewClient()
	helper, err := auth.NewHelper(client)
	if err != nil {
		log.Fatalf("Failed to create helper: %v", err)
	}

	// Function to handle token refresh
	var mu sync.Mutex
	tokenRefreshChan := make(chan struct{})

	go func() {
		for {
			select {
			case <-tokenRefreshChan:
				// Reset the timer, continue to the next iteration
				continue
			case <-time.After(23*time.Hour + 59*time.Second):
				mu.Lock()
				_, err := helper.GetToken()
				if err != nil {
					log.Printf("Failed to refresh token: %v", err)
				} else {
					log.Println("Token refreshed successfully")
				}
				mu.Unlock()
			}
		}
	}()

	// Simulate an API call
	for {
		token, err := helper.GetToken()
		if err != nil {
			log.Fatalf("Failed to get token: %v", err)
		}

		fmt.Printf("Access token: %s\n", token)

		// Signal the refresh routine to reset the timer
		tokenRefreshChan <- struct{}{}

		// Simulate waiting before the next API call
		time.Sleep(10 * time.Minute)
	}
}
