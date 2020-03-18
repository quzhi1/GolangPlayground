package helper

import (
	"fmt"
	"sync"
)

// Fetcher returns the body of URL and
// a slice of URLs found on that page.
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

// urlMap is a cache for crawler
type cache struct {
	visited map[string]bool
	mux     sync.Mutex
}

func (c *cache) checkVisitied(url string) bool {
	c.mux.Lock()
	defer c.mux.Unlock()

	_, ok := c.visited[url]
	if ok {
		return true
	}
	c.visited[url] = true
	return false
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	crawlerCache := &cache{visited: map[string]bool{}}
	var wg sync.WaitGroup
	wg.Add(1)
	go recursiveCrawl(url, depth, fetcher, crawlerCache, &wg)
	wg.Wait()
	// time.Sleep(5 * time.Second)
}

func recursiveCrawl(url string, depth int, fetcher Fetcher, crawlerCache *cache, wg *sync.WaitGroup) {
	defer wg.Done()
	if depth <= 0 {
		// close(resultChan)
		return
	}

	// Synchronize
	if crawlerCache.checkVisitied(url) {
		return
	}

	// Do crawling
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go recursiveCrawl(u, depth-1, fetcher, crawlerCache, wg)
	}
	return
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

// TestWebCrawler verifies Crawl function
func TestWebCrawler() {
	Crawl("https://golang.org/", 4, fetcher)
}
