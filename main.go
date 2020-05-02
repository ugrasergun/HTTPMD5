package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {

	threadCount := flag.Int("parallel", 10, "Thread Count")

	flag.Parse()

	fmt.Println(*threadCount)

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
			time.Sleep(2 * time.Second)
		}(url)
	}

	wg.Wait()
}
