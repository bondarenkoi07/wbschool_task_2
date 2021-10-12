package main

import "regexp"

var (
	regex = regexp.MustCompile(`https?:\\/\\/(www\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\\+.~#?&/=]*)`)
)

func Download(link string) {

}

func main() {

}
