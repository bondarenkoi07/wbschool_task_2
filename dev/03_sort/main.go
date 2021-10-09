package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
//InvalidColumnError = errors.New("column includes non-numeric values")
)

func main() {
	var (
		reader        *os.File
		err           error
		numeric       = flag.Bool("n", false, "сортировать по числовому значению")
		column        = flag.Int("k", 1, "указание колонки для сортировки")
		reverse       = flag.Bool("r", false, "сортировать в обратном порядке")
		unique        = flag.Bool("u", false, "не выводить повторяющиеся строки")
		checkIsSorted = flag.Bool("c", false, "проверять отсортированы ли данные")
		filepath      = flag.String("filepath", "", "путь к сортируемому файлу")
	)

	flag.Parse()

	if *filepath == "" {
		reader = os.Stdin
	} else {
		reader, err = os.Open(*filepath)
		if err != nil {
			log.Fatal(err)
		}
	}

	buf := bufio.NewReader(reader)
	//log.Printf("fetching strings... buffer size: %d", buf.Size())

	fetchedStrings, err := FetchStrings(buf)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return
	}

	//log.Printf("fetched strings: %v\n", fetchedStrings)

	if *unique {
		fetchedStrings = UniqueStrings(fetchedStrings)
		log.Printf("uniqe fetched strings: %v\n", fetchedStrings)
	}

	var unpackedStrings = UnpackStrings(fetchedStrings)
	//log.Printf("unpacked fetched strings: %v\n", fetchedStrings)
	if *checkIsSorted {
		var sorted bool
		sorted, err = SortingAction(unpackedStrings, *column-1, *numeric, *reverse, sort.SliceIsSorted)
		if err != nil {
			fmt.Println("sort error: ", err)
			return
		}

		if sorted {
			fmt.Println("source is already sorted")
		} else {
			fmt.Println("source is not" +
				" already sorted")
		}
		return
	}

	_, err = SortingAction(
		unpackedStrings,
		*column-1,
		*numeric,
		*reverse,
		func(x interface{}, less func(i int, j int) bool) bool {
			sort.Slice(x, less)
			return true
		},
	)
	if err != nil {
		fmt.Println("sort error: ", err)
		return
	}

	_, err = sortPacker(unpackedStrings).WriteTo(os.Stdout)
	if err != nil {
		fmt.Printf("bad syntax: %v", err)
		return
	}

}

func FetchStrings(buf *bufio.Reader) ([]string, error) {
	var (
		output     = make([]string, 0, 0)
		err        error
		readString string
	)

	for err != io.EOF {
		readString, err = buf.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Printf("this error: %v\n", err)
			return nil, err
		}
		//fmt.Printf("fetched! %s\n", readString)
		output = append(output, readString)
	}

	return output, nil
}

func UnpackStrings(input []string) [][]string {
	var (
		sorting [][]string
		index   int
	)
	for index = 0; index < len(input)-1; index++ {
		sorting = append(sorting, strings.Fields(input[index]))
	}

	if input[index] != "" {
		sorting = append(sorting, strings.Fields(input[index]))
	}

	return sorting
}

func SortingAction(
	input [][]string,
	col int,
	numeric bool,
	reverse bool,
	sorter func(x interface{}, less func(i int, j int) bool) bool,
) (bool, error) {
	var (
		err error
	)
	if len(input) == 0 {
		return false, errors.New("empty source")
	}
	if col > len(input[0]) || col < 0 {
		return false, errors.New(fmt.Sprintf("no such column %d", col))
	}
	var less func(i, j int) bool
	if numeric {

		less = func(i, j int) bool {

			if err != nil {
				return false
			}

			var valI, valJ int
			valI, err = strconv.Atoi(input[i][col])
			if err != nil {
				return false
			}

			valJ, err = strconv.Atoi(input[j][col])
			if err != nil {
				return false
			}

			if reverse {
				return !(valI < valJ)
			}

			return valI < valJ
		}
	} else {
		less = func(i, j int) bool {
			if reverse {
				return !(input[i][col] < input[j][col])
			}
			return input[i][col] < input[j][col]
		}
	}

	sorter(input, less)

	return sorter(input, less), err
}

func sortPacker(input [][]string) *strings.Reader {
	var output []string
	for index := range input {
		output = append(output, strings.Join(input[index], " "))
	}

	return strings.NewReader(strings.Join(output, "\n"))
}

func UniqueStrings(input []string) []string {
	var (
		output  []string
		uniquer = make(map[string]struct{})
	)
	for index := range input {
		uniquer[input[index]] = struct{}{}
	}

	for index := range uniquer {
		output = append(output, index)
	}

	return output
}
