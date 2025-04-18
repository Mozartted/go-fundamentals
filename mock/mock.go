package mock

import (
	"fmt"
	"io"
	"iter"
	"time"
)

const finalWord = "Go!"

// this is the interface to mock
type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func countDownFrom(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for num := range countDownFrom(3) {
		fmt.Fprint(out, num)
		sleeper.Sleep()
	}

	// iteratorflow := countDownFrom(3)

	// iteratorflow(func (i int){
	// 	fmt.Fprint(out, i)
	// })

	// fmt.Fprint(out, finalWord)
}
