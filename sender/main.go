package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	numRequests = 10000
	url         = "http://localhost:8080/send"
)

func main() {
	transport := &http.Transport{
		MaxIdleConns:        10, // Adjust the maximum number of idle connections
		MaxIdleConnsPerHost: 10, // Adjust the maximum number of idle connections per host
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}

	var wg sync.WaitGroup
	requestsCh := make(chan int, numRequests)

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sendRequest(client, i)
		}(i)
		requestsCh <- i
	}

	close(requestsCh)
	wg.Wait()
}

func sendRequest(client *http.Client, i int) {
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
		fmt.Printf("Request error for request %d: %v\n", i, err)
		return
	}

	defer resp.Body.Close()

	fmt.Printf("Response status code for request %d: %d\n", i, resp.StatusCode)
}
