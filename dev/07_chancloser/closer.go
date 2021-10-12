package _7_chancloser

import (
	"log"
	"sync"
)

func Closer(cs ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	closer := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for index, c := range cs {
		go func(c <-chan interface{}, index int) {
			defer func() {
				wg.Done()
			}()
			for {
				select {
				case <-closer:
					c = nil
					return
				case val, ok := <-c:
					if !ok {
						for i := 0; i < len(cs)-1; i++ {
							closer <- struct{}{}
						}
						return
					}
					out <- val
				}
			}
		}(c, index)
	}
	wg.Wait()
	log.Println("Here")
	close(out)
	return out
}
