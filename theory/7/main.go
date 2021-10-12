package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			// при чтении из закрытого канала v будет принимать дефолтные значения для типа каналов a
			//потому что не происходить проверки на закрытие канала ( v, ok := <- a). И если оба канала закрыты,
			//то прервать цикл и закрыть с
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}

		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v) // выведет с 1 по 8, после чего будет слать 0 до бесконечности
	}
}
