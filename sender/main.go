package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

func main() {
	numRequests := 100
	url := "http://localhost:8080/send"

	var wg sync.WaitGroup
	wg.Add(numRequests)

	requestsCh := make(chan int, numRequests)
	for i := 0; i < numRequests; i++ {
		go func(i int) {
			defer wg.Done()

			client := &http.Client{}

			// Dynamically generate the value for the "message" key
			value := generateRandomString()

			jsonData, err := json.Marshal(value)
			if err != nil {
				fmt.Println("Error creating JSON data:", err)
				return
			}

			req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
			if err != nil {
				fmt.Println("Error creating request:", err)
				return
			}

			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Request error:", err)
				return
			}
			defer resp.Body.Close()

			fmt.Printf("Response status code for request %d: %d\n", i, resp.StatusCode)
		}(i)

		requestsCh <- i
	}

	close(requestsCh)
	wg.Wait()
}
