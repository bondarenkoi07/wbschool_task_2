package _7_chancloser

import (
	"fmt"
	"testing"
	"time"
)

func TestCloser(t *testing.T) {
	var or = Closer
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	end := time.Since(start)
	fmt.Printf(`fone after %v`, end)
	if (2 * time.Second) < end {
		t.Error("time exceeded,  something goes wrong(")
	}
}
