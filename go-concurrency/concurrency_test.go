package concurrency

import (
	"reflect"
	"testing"
	"time"
	"github.com/aszenz/basic-golang/go-test-helpers"
	"net/http/httptest"
	"net/http"
)

// ---------- //
func slowStubWebsiteChecker(_ string) bool{
	time.Sleep(20 * time.Millisecond)
	return true
}

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
        "http://google.com",
        "http://blog.gypsydave5.com",
        "waat://furhurterwe.geds",
	}

	want := map[string]bool{
			"http://google.com":          true,
			"http://blog.gypsydave5.com": true,
			"waat://furhurterwe.geds":    false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i:= 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

func makeDelayedHTTPServer(delay time.Duration) *httptest.Server {
	 return httptest.NewServer(http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
// ---------- //
func TestWebsiteRacer(t *testing.T) {
	t.Run("Returns the faster of the two urls", func(t *testing.T) {
		slowServer := makeDelayedHTTPServer(20 * time.Millisecond)	
		fastServer := makeDelayedHTTPServer(10 * time.Millisecond)	

		defer slowServer.Close()
		defer fastServer.Close()

		got, err := WebsiteRacer(slowServer.URL, fastServer.URL, 50*time.Millisecond)
		want := fastServer.URL

		testhelpers.AssertNoError(t, err)
		testhelpers.AssertStringsEqual(t, got, want)
	})
	t.Run("Return an error if server does not respond within 10 secs", func(t *testing.T){
		serverA := makeDelayedHTTPServer(11 * time.Millisecond)
    serverB := makeDelayedHTTPServer(12 * time.Millisecond)

    defer serverA.Close()
    defer serverB.Close()

		_, err := WebsiteRacer(serverA.URL, serverB.URL, 10*time.Millisecond)

		if err == nil {
        t.Error("expected an error but didn't get one")
    }
	})
}