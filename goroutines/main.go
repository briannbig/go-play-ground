package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	respCh := make(chan any, 2)

	start := time.Now()
	username := fetchUser()

	var wg sync.WaitGroup
	wg.Add(2)

	go fetchUserLikes(username, respCh, &wg)
	go fetchUserMatch(username, respCh, &wg)

	wg.Wait()
	close(respCh)

	for resp := range respCh {
		fmt.Println("resp: ", resp)
	}

	fmt.Println("took: ", time.Since(start))

}

func fetchUser() string {
	time.Sleep(100 * time.Millisecond)

	return "jdoe"
}

func fetchUserLikes(username string, respCh chan any, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("fetching likes for", username)
	time.Sleep(150 * time.Millisecond)

	respCh <- 11
}

func fetchUserMatch(username string, ch chan any, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("fetching match for", username)
	time.Sleep(125 * time.Millisecond)

	ch <- "jane"
}
