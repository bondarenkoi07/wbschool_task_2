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
				log.Printf("chan %d  finally closed", index)
			}()
			log.Printf("chan %d started\n", index)
			for {
				select {
				case <-closer:
					log.Printf("chan %d received closer\n", index)
					return

				case val, ok := <-c:
					log.Printf("chan %d send val\n", index)
					if !ok {
						log.Printf("chan %d  closed\n", index)
						for i := 0; i < len(cs)-1; i++ {
							closer <- struct{}{}
						}

						log.Printf("goroutine with first closed chan %d\n", index)

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
