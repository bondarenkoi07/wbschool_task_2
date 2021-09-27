package main

import (
	"fmt"
)

func NewDict(input []string) map[string][]string {
	var output = make(map[string][]string)
	var sets = make(map[string]map[rune]struct{})
	for _, inputString := range input {
		_, isSet := sets[inputString]
		fmt.Printf("%s isSet? %v\n", inputString, isSet)
		if !isSet {
			check := false
			for i, runeSet := range sets {
				if len([]rune(i)) != len([]rune(inputString)) {
					fmt.Printf("set %v string %v |%v versus %v\n", i, inputString, len(runeSet), len([]rune(inputString)))
				} else {
					for _, elems := range inputString {
						_, isElemSet := runeSet[elems]
						if !isElemSet {
							check = false
							break
						} else {
							check = true
						}
					}
					if check {
						output[i] = append(output[i], inputString)
						break
					}
				}
			}

			if !check {
				output[inputString] = make([]string, 0, len(input))
				set := make(map[rune]struct{})
				for _, r := range inputString {
					set[r] = struct{}{}
				}
				sets[inputString] = set
			}
		}
	}

	for index, set := range output {
		if len(set) == 0 {
			delete(output, index)
		}
	}

	return output
}

func main() {
	var ins = []string{
		"пятак", "тяпка", "стол", "слот", "пятка", "гусь", "тпкяа", "усь", "тпка",
	}

	out := NewDict(ins)
	fmt.Println(out)
}
