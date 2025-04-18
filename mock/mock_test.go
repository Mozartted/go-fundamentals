package mock

import (
	"bytes"
	"testing"
)

// injecting the SpySleeper during the test.

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeper := SpySleeper{}
	Countdown(buffer, &sleeper)

	got := buffer.String()
	want := `321Go!`

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
