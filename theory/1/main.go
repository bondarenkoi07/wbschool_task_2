package main

import (
	"fmt"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] // создается срез массива со второго по четвертый элемента
	fmt.Println(b)       // [77 78 79] выводит со второго по четвертый элементы массива
}