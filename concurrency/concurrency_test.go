package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestConcurrency(t *testing.T) {
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

	results := CheckWebsite(mockWebChecker, websites)

	if !reflect.DeepEqual(results, want) {
		t.Fatalf("wanted %v, got %v", want, results)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsite(slowStubWebsiteChecker, urls)
	}
}
