package main

import (
	"io"
	"regexp"
	"strings"
	"testing"
)

var (
	builder strings.Builder
)

func TestInit(t *testing.T) {
	match := &MatchedFlag{builder: &builder}
	post := &PostFlag{builder: &builder, buf: make([][]byte, 0, 2)}
	pre := &PreFlag{builder: &builder, matched: match, buf: make([][]byte, 0, 2), post: post}
	counter := &CounterFlag{next: pre, builder: &builder}

	reg := regexp.MustCompile("hol")

	grep := Grep{builder: &builder, regex: reg, output: counter}
	reader := strings.NewReader("hello\nauf\nfoobar\nholla\njust\nmake\nholla\nhol\nfoooooo\nheheheh")
	err := grep.DoStaff(reader)
	if err != nil && err != io.EOF {
		t.Fatal(err)
	}
}
