package main

import (
	"net/http"
	"sync"
	"fmt"
	"time"
)

// Add comma separated your website's urls
// Note even the last one in the list must include a comma in the end
var urls = []string {
	// "https://www.example.com",
	// "https://www.example2.com",
}

func getStatus(url string, wg *sync.WaitGroup) (string, error) {
	req, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	wg.Done()
	fmt.Printf("url: %s Status Code: %s\n", url, req.Status)
	return req.Status, nil
}

func printStatus() {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go getStatus(url, &wg)
	}
	wg.Wait()
}

func main() {
	// be careful to not DDoS your pages
	const SleepTime = time.Second * 60

	for {
		printStatus()
		time.Sleep(SleepTime)
	}
}
