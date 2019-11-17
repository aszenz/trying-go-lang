package concurrency

import (
	"fmt"
	"net/http"
	"time"
)
// WebsiteChecker checks if a website is available  
type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// CheckWebsites uses websiteChecker to check each website and returns a map
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}
	for i:=0; i< len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}
	return results
}

func measureHTTPResponseTime(url string) time.Duration {
	startA := time.Now()
	http.Get(url)
	return time.Since(startA)
}

func inefficientWebsiteRacer(url1 string, url2 string) string {
	url1Duration := measureHTTPResponseTime(url1)
	url2Duration := measureHTTPResponseTime(url2)

	if url1Duration < url2Duration {
		return url1
	}

	return url2
}

func ping(url string) chan bool{
	ch := make(chan bool)
	go func(){
		http.Get(url)
		ch <- true
	}()
	return ch
}

// WebsiteRacer returns the faster of the two urls
func WebsiteRacer(url1 string, url2 string, timeout time.Duration) (string, error) {
	select {
	case <- ping(url1):
		return url1, nil
	case <- ping(url2):
		return url2, nil
	case <- time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", url1, url2)
	}
}
