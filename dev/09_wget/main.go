package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var (
	regex = regexp.MustCompile(`https?:\\/\\/(www\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\\+.~#?&/=]*)`)
)

func Download(link string) error {
	resp, err := http.Get(link)
	if err != nil {
		log.Println("connection")
		return err
	}

	var b bytes.Buffer
	var builder strings.Builder
	_, err = io.Copy(&b, resp.Body)

	if err != nil {
		log.Println("connection")
		return err
	}
	var readString string
	for err != io.EOF {
		readString, err = b.ReadString('\n')
		builder.Grow(len(readString))
		builder.WriteString(readString)
		if err == io.EOF {
			log.Println("EOF")
			break
		} else if err != nil {
			log.Println("reading")
			return err
		}

		lnk := regex.Find([]byte(readString))
		if len(lnk) > 0 {
			recErr := Download(string(lnk))
			if recErr != nil {
				log.Println(recErr)
			}
		}
	}
	log.Println(removeSlash(link))
	file, err := os.Create(removeSlash(link))
	if err != nil {
		log.Println("create")
		return err
	}
	defer file.Close()

	defer resp.Body.Close()

	written, err := io.Copy(file, strings.NewReader(builder.String()))
	log.Println("copy ", written)
	return err
}

func removeSlash(input string) string {
	var output strings.Builder
	output.Grow(len(input))
	for _, i2 := range input {
		if i2 != '/' {
			output.WriteRune(i2)
		} else {
			output.WriteRune('_')
		}
	}

	return output.String()
}

func main() {
	var link = flag.String("host", "http://db1.mati.su", "host to download")
	flag.Parse()
	err := Download(*link)
	if err != nil {
		log.Println(err)
	}
}
