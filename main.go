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

	urls := flag.Args()

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			uri, md5 := getURIMD5(url)
			fmt.Println(uri, md5)
			time.Sleep(2 * time.Second)
			defer wg.Done()
		}(url)
	}

	wg.Wait()
}
