package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func StringUnpack(input string) string {
	var (
		output strings.Builder
		tmp    = make([]rune, 0, len(input))
		index  = -1
		val    rune
	)

	output.Grow(len(input))

	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic occurred :%v\n at index: %d and rune is  %c", err, index, val)
			log.Printf("what was built: %v\n", output.String())
		}
	}()

	for _, val = range input {
		index++
		if index < 0 {
			continue
		}
		if unicode.IsDigit(val) {
			var i int
			var r rune
			if len(tmp)-1 < 0 {
				fmt.Println("incorrect string")
				return ""
			}
			doppler := tmp[len(tmp)-1]
			num, _ := strconv.Atoi(string(val))
			output.WriteString(string(tmp))
			tmp = make([]rune, 0, len(input))
			skip := -1
			for i, r = range input[index+1:] {
				//a4bc2d5e
				if !unicode.IsDigit(r) {
					break
				} else {
					skip--
					tmpNum, _ := strconv.Atoi(string(r))
					num = 10*num + tmpNum
				}
			}
			i = index + i
			input = input[i+1:]
			fmt.Printf("input is %s and index is %d, num is %d\n", input, index, num)
			index = skip
			for i = 0; i < num-1; i++ {
				output.WriteRune(doppler)
			}
		} else {
			tmp = append(tmp, val)
		}
	}
	output.WriteString(string(tmp))

	return output.String()
}

func main() {

	fmt.Println("Enter packed string:")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err)
	}

	StringUnpack(input)

	/*
		for i := 0;i<len(input);i++ {
			if unicode.IsDigit(rune(input[i])){
			}else{
				tmp = append(tmp, rune(input[i]))
			}
		}
	*/
}
