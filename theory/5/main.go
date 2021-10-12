package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil // value == nil, dynamic type != *customError
}

func main() {
	var err error
	err = test()
	if err != nil { // value == nil, dynamic type != nil
		println("error")
		return
	}
	println("ok")
}
