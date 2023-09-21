//////////////////////////////////////////////////////////////////////
//
// Your task is to change the code to limit the crawler to at most one
// page per second, while maintaining concurrency (in other words,
// Crawl() must be called concurrently)
//
// @hint: you can achieve this by adding 3 lines
//

package main

import (
	"fmt"
	"sync"
	"time"
)

type CrawlScheduler struct {
	wg       sync.WaitGroup
	throttle <-chan time.Time
}

func NewCrawlScheduler() *CrawlScheduler {
	return &CrawlScheduler{
		wg: sync.WaitGroup{},
		// 1 page per second
		throttle: time.Tick(1 * time.Second),
	}
}

// Crawl uses `fetcher` from the `mockfetcher.go` file to imitate a
// real crawler. It crawls until the maximum depth has reached.
func (cs *CrawlScheduler) Crawl(url string, depth int) {
	defer cs.wg.Done()

	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s found: %s %q\n", time.Now().String(), url, body)

	cs.wg.Add(len(urls))
	for _, u := range urls {
		// Do not remove the `go` keyword, as Crawl() must be
		// called concurrently
		<-cs.throttle
		go cs.Crawl(u, depth-1)
	}
	return
}

func main() {
	c := NewCrawlScheduler()
	c.wg.Add(1)

	c.Crawl("http://golang.org/", 4)
	c.wg.Wait()
}
