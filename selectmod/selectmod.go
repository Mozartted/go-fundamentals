package selectmod

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(slowURL string, fastUrl string) string {
	duration := measureURLTime(slowURL)
	durationB := measureURLTime(fastUrl)

	result := fmt.Sprintf("Duration 1: %d , Duration 2: %d", duration, durationB)

	fmt.Println(result)
	if duration > durationB {
		return fastUrl
	} else {
		return slowURL
	}
}

func measureURLTime(slowURL string) time.Duration {
	startTimeA := time.Now()

	http.Get(slowURL)
	return time.Since(startTimeA)
}
