package main

import (
	"io"
	"regexp"
	"strings"
)

type IFlagChainNode interface {
	SetBuilder(*strings.Builder)
	DoStaff(bool)
	Next(bool) *IFlagChainNode
}

type Grep struct {
	reader io.Reader
	output strings.Builder
	regex  regexp.Regexp
	flag   *IFlagChainNode
}

func (g Grep) Match() {

}
