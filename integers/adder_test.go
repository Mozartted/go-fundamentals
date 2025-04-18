package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	defer fmt.Println("end of call")
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Sum %q is not equal to Expected %q", sum, expected)
	}
}
