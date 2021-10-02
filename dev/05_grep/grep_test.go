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
	pre := &PreFlag{builder: &builder, buf: make([][]byte, 0, 2), next: match}
	post := &PostFlag{builder: &builder, buf: make([][]byte, 0, 2), next: pre}
	counter := &CounterFlag{next: post, builder: &builder}

	reg := regexp.MustCompile("hol")

	grep := Grep{builder: &builder, regex: reg, output: counter}
	reader := strings.NewReader("hello\nauf\nfoobar\nholla\njust\nmake\nholla\nhol\nfoooooo\nheheheh")
	err, _ := grep.DoStaff(reader)
	if err != nil && err != io.EOF {
		t.Fatal(err)
	}
}

func TestInit2(t *testing.T) {
	match := &MatchedFlag{builder: &builder}
	pre := &PreFlag{builder: &builder, buf: make([][]byte, 0, 2), next: match}
	post := &PostFlag{builder: &builder, buf: make([][]byte, 0, 2), next: pre}
	counter := &CounterFlag{next: post, builder: &builder}

	reg := regexp.MustCompile("foo")

	grep := Grep{builder: &builder, regex: reg, output: counter}
	reader := strings.NewReader("hello\nauf\nfoobar\nholla\njust\nmake\nholla\nhol\nfoooooo\nheheheh")
	err, _ := grep.DoStaff(reader)
	if err != nil && err != io.EOF {
		t.Fatal(err)
	}
}
