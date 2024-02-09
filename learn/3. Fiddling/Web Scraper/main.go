package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Worker function that fetches the HTML content of a URL and sends it through a channel
func fetch(url string, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to fetch %s: %s\n", url, resp.Status)
		return
	}

	// Read the body of the response
	body := make([]byte, 0)
	for {
		buffer := make([]byte, 1024)
		n, err := resp.Body.Read(buffer)
		if n == 0 || err != nil {
			break
		}
		body = append(body, buffer[:n]...)
	}

	// Send the body content through the channel
	ch <- string(body)
}

func main() {
	urls := []string{"https://www.facebook.com/", "https://www.google.com", "https://www.github.com"}

	// Create a channel to receive HTML content
	ch := make(chan string)

	var wg sync.WaitGroup

	// Spawn a Goroutine for each URL
	for _, url := range urls {
		wg.Add(1)
		go fetch(url, &wg, ch)
	}

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(ch) // Close the channel once all workers are done
	}()

	// Receive and print HTML content from the channel
	for html := range ch {
		fmt.Println(html[:100]) // Print only the first 100 characters of each HTML content
	}
}
