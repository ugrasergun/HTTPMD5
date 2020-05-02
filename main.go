package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {

	threadCount := flag.Int("parallel", 10, "Thread Count")

	flag.Parse()

	semaphore := make(chan bool, *threadCount)

	urls := flag.Args()

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		semaphore <- true

		go func(url string) {
			defer wg.Done()
			defer func() { <-semaphore }()
			uri, md5 := getURIMD5(url)
			fmt.Println(uri, md5)
		}(url)
	}

	wg.Wait()
}
