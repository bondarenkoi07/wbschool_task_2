package main

import (
	"fmt"
)

func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return // именнованое возвращаемое значение будет изменено defer'ом.
}

func anotherTest() int {
	var x int
	defer func() {
		x++
	}() // Выполнится после return, поэтому не повлияет на возвращаемое значение
	x = 1
	return x
}

func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
