package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var (
		reader     *os.File
		err        error
		lines      []string
		cols       []int
		filepath   = flag.String("filepath", "", "путь до источника")
		separator  = flag.String("d", "\t", "использовать другой разделитель")
		rawColumns = flag.String("f", "", "выбрать поля (колонки)")
		separated  = flag.Bool("s", false, "только строки с разделителем")
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
	lines, err = FetchStrings(buf)

	if *rawColumns == "" {
		fmt.Println(
			"cut: вы должны задать список байт, символов или полей\n" +
				"По команде «cut --help» можно получить дополнительную информацию.",
		)
		return
	}

	cols, err = ParseColumns(*rawColumns)

	if err != nil {
		fmt.Println(err)
		return
	}

	cut, err := FieldCut(lines, *separator, *separated, cols...)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cut)
}

func ParseColumns(rawCols string) ([]int, error) {
	atoi, err := strconv.Atoi(rawCols)
	if err == nil {
		return []int{atoi - 1}, nil
	}

	regex := regexp.MustCompile("^(\\d+)-(\\d+)$")
	parsedCols := regex.FindSubmatch([]byte(rawCols))

	if len(parsedCols) == 3 {
		a, _ := strconv.Atoi(string(parsedCols[1]))
		b, _ := strconv.Atoi(string(parsedCols[2]))

		if a > b {
			a, b = b, a
		}

		output := make([]int, 0, b-a+1)
		for i := a - 1; i < b; i++ {
			output = append(output, i)
		}
		return output, nil
	}

	parsedCommaCols := strings.Split(rawCols, ",")

	if len(parsedCommaCols) != 0 {
		output := make([]int, 0, len(parsedCommaCols))
		col := 0
		for _, value := range parsedCommaCols {
			col, err = strconv.Atoi(value)
			if err != nil {
				break
			}
			output = append(output, col-1)
		}

		return output, nil
	}
	err = errors.New("cut: wrong columns format: cols must be separated by '-' or ',' or be single")
	return nil, err
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
			return nil, err
		}
		//fmt.Printf("fetched! %s\n", readString)
		output = append(output, readString)
	}

	return output, nil
}

func FieldCut(input []string, sep string, separated bool, colNum ...int) (string, error) {
	var output = strings.Builder{}
	for _, value := range input {
		var builder = strings.Builder{}
		builder.Grow(len(value))

		strSplit := strings.Split(value, sep)

		if separated && len(strSplit) < 1 {
			continue
		}

		for _, index := range colNum {
			if index >= len(strSplit) {
				continue
				// return nil, errors.New(fmt.Sprintf("invalid column, got %d max %d", index,len(strSplit)))
			}
			_, err := builder.WriteString(strSplit[index])
			if err != nil {
				return "", err
			}
			_, err = builder.WriteRune(' ')
			if err != nil {
				return "", err
			}
		}
		log.Printf("was built %v\n", builder.String())
		output.Grow(builder.Len() + 1)
		output.WriteString(builder.String())
		output.WriteRune('\n')
	}

	return output.String(), nil
}
