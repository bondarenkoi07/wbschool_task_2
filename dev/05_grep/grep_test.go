package main

import (
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
	"testing"
)

var (
	builder = &strings.Builder{}
	match   = &MatchedFlag{builder: builder}
	pre     = &PreFlag{builder: builder, buf: make([][]byte, 0, 1), next: match}
	post    = &PostFlag{builder: builder, buf: make([][]byte, 0, 1), next: pre}
	counter = &CounterFlag{builder: builder}
	index   = &IndexFlag{next: post}
	node    IFlagChainNode

	readerDefault              = &strings.Reader{}
	reader                     = "hello\nauf\nfoobar\nholla\njust\nmake\nholla\nhol\nfoooooo\nheheheh"
	readerDefaultOutput        = "holla\nholla\nhol\n"
	readerDefaultNumOutput     = "4:holla\n7:holla\n8:hol\n"
	readerDefaultPostOutput    = "holla\njust\nholla\nhol\nfoooooo\n"
	readerDefaultPreOutput     = "foobar\nholla\nmake\nholla\nhol\n"
	readerDefaultContextOutput = "foobar\nholla\njust\nmake\nholla\nhol\nfoooooo\n"
	readerDefaultCountOutput   = "3"
	readerEmpty                = strings.NewReader("")
)

func TestDefaultMatching(t *testing.T) {
	readerDefault = strings.NewReader(reader)
	node = match
	reg := regexp.MustCompile("hol")

	grep := Grep{builder: builder, regex: reg, output: node}
	err, output := grep.DoStaff(readerDefault)
	if err != nil && err != io.EOF {
		t.Fatal(err)
	}

	if output != readerDefaultOutput {
		t.Fatalf("invalid output: %v and %v", output, readerDefaultOutput)
	}

	builder.Reset()
	readerDefault = strings.NewReader(reader)
}

func TestEmptyInput(t *testing.T) {
	node = match
	reg := regexp.MustCompile("hol")

	grep := Grep{builder: builder, regex: reg, output: node}
	fmt.Println(readerDefault)
	err, output := grep.DoStaff(readerEmpty)
	if err != nil && err != io.EOF {
		t.Fatal(err)
	}

	if output != "" {
		t.Fatalf("invalid output: %v and %v", output, "(empty)")
	}
	builder.Reset()
}

func TestNum(t *testing.T) {
	node = index
	node.SetNext(match)
	reg := regexp.MustCompile("hol")
	log.Println(node)
	grep := Grep{builder: builder, regex: reg, output: node}

	err, output := grep.DoStaff(readerDefault)
	if err != nil && err != io.EOF {
		t.Fatal(err)
	}

	if output != readerDefaultNumOutput {
		t.Fatalf("invalid output: %v and %v", output, readerDefaultNumOutput)
	}
	builder.Reset()
	readerDefault = strings.NewReader(reader)
}

func TestAfter(t *testing.T) {
	node = post
	node.SetNext(match)
	reg := regexp.MustCompile("hol")

	grep := Grep{builder: builder, regex: reg, output: node}

	err, output := grep.DoStaff(readerDefault)
	if err != nil && err != io.EOF {
		t.Fatal(err)
	}

	if output != readerDefaultPostOutput {
		t.Fatalf("invalid output: %v and %v", output, readerDefaultPostOutput)
	}
	builder.Reset()
	readerDefault = strings.NewReader(reader)
}

func TestBefore(t *testing.T) {
	node = pre
	node.SetNext(match)
	reg := regexp.MustCompile("hol")

	grep := Grep{builder: builder, regex: reg, output: node}

	err, output := grep.DoStaff(readerDefault)
	if err != nil && err != io.EOF {
		t.Fatal(err)
	}

	if output != readerDefaultPreOutput {
		t.Fatalf("invalid output: %v and %v", output, readerDefaultPreOutput)
	}
	builder.Reset()
	readerDefault = strings.NewReader(reader)
}

func TestContext(t *testing.T) {
	node = post
	node.SetNext(pre)
	pre.SetNext(match)
	reg := regexp.MustCompile("hol")

	grep := Grep{builder: builder, regex: reg, output: node}

	err, output := grep.DoStaff(readerDefault)
	if err != nil && err != io.EOF {
		t.Fatal(err)
	}

	if output != readerDefaultContextOutput {
		t.Fatalf("invalid output: %v and %v", output, readerDefaultContextOutput)
	}
	builder.Reset()
	readerDefault = strings.NewReader(reader)
}

func TestCount(t *testing.T) {
	node = counter
	reg := regexp.MustCompile("hol")

	grep := Grep{builder: builder, regex: reg, output: node}

	err, output := grep.DoStaff(readerDefault)
	if err != nil && err != io.EOF {
		t.Fatal(err)
	}

	if output != readerDefaultCountOutput {
		t.Fatalf("invalid output: %v and %v", output, readerDefaultCountOutput)
	}
	builder.Reset()
	readerDefault = strings.NewReader(reader)
}
