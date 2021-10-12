package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err // value nil, dynamic type *os.PathError
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil) // for nil - value nil, dynamic type nil
}
