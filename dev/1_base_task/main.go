package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(time)
}
