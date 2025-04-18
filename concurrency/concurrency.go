package concurrency

import (
	"sync"
)

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsite(wc WebsiteChecker, urls []string) map[string]bool {
	m := sync.RWMutex{}
	results := make(map[string]bool)

	chanResults := make(chan result)

	for _, url := range urls {
		go func() {
			m.Lock()
			chanResults <- result{url, wc(url)}
			m.Unlock()
		}()
	}

	for i := 0; i < len(urls); i++ {
		r := <-chanResults
		results[r.string] = r.bool
	}

	return results
}
