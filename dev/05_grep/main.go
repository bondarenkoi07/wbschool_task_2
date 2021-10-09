package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	var (
		numCounterPtr = flag.Bool("n", false, "печатать номер строки")
		CounterPtr    = flag.Bool("c", false, "Вывод количества подходящих строк")
		AfterPtr      = flag.Int("A", -1, "Вывод n строк после совпадения")
		BeforePtr     = flag.Int("B", -1, "Вывод n строк до совпадения")
		ContextPtr    = flag.Int("C", -1, "Вывод n строк около совпадения")
		ignoreCasePtr = flag.Bool("i", false, "игнорировать регистр")
		FixedPtr      = flag.Bool("F", false, "Вывод количества полностью совпадающих строк строк")
		inversionPtr  = flag.Bool("v", false, "вывод несовпадений")
		builder       = &strings.Builder{}
		pattern       = flag.String("pattern", "", "паттерн поиска")
		filepath      = flag.String("filepath", "", "путь к файлу")
		reader        *os.File
		node          IFlagChainNode
		root          IFlagChainNode
	)

	flag.Parse()

	if *filepath == "" {
		reader = os.Stdin
	} else {
		file, err := os.OpenFile(*filepath, os.O_RDONLY, 0744)
		if err != nil {
			log.Fatalf("error: %v, filepath %s", err, *filepath)
		}

		reader = file
	}

	if *FixedPtr {
		*pattern = fmt.Sprintf("^%s(\\r\\n|\\r|\\n)$", *pattern)
	}

	if *ignoreCasePtr {
		*pattern = fmt.Sprintf("(?i)%s", *pattern)
	}

	regex := regexp.MustCompile(*pattern)

	var flags = make([]IFlagChainNode, 0, 5)

	if *CounterPtr {
		flags = append(flags, &CounterFlag{builder: builder})
	} else {
		if *numCounterPtr {
			flags = append(flags, &IndexFlag{})
		}

		if *ContextPtr > 0 {
			flags = append(flags, &PostFlag{builder: builder, buf: make([][]byte, 0, *ContextPtr)})
			flags = append(flags, &PreFlag{builder: builder, buf: make([][]byte, 0, *ContextPtr)})
		} else {
			if *AfterPtr > 0 {
				flags = append(flags, &PostFlag{builder: builder, buf: make([][]byte, 0, *AfterPtr)})
			}
			if *BeforePtr > 0 {
				flags = append(flags, &PreFlag{builder: builder, buf: make([][]byte, 0, *BeforePtr)})
			}
		}
	}
	node = nil
	for _, chainNode := range flags {
		if node == nil {
			node = chainNode
			root = chainNode
		} else {
			node.SetNext(chainNode)
			node = chainNode
		}
	}
	if node == nil {
		tmp := &MatchedFlag{builder: builder}
		node = tmp
		root = tmp
	} else {
		node.SetNext(&MatchedFlag{builder: builder})
	}

	grepMaster := Grep{builder: builder, regex: regex, output: root, exclude: *inversionPtr}

	err, str := grepMaster.DoStaff(reader)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Println(str)
}
