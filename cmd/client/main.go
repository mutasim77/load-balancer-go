package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			resp, err := http.Get("http://localhost:8080")
			if err != nil {
				fmt.Printf("Request %d failed: %v\n", id, err)
				return
			}
			defer resp.Body.Close()
			fmt.Printf("Request %d: Status %s\n", id, resp.Status)
		}(i)
		time.Sleep(100 * time.Millisecond)
	}
	wg.Wait()
}
