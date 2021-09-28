package main

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

//

type IFlagChainNode interface {
	SetBuilder(*strings.Builder)
	DoStaff(bool, []byte)
	Next(bool) *IFlagChainNode
}

type PrevFlag struct {
	builder *strings.Builder
}

func (p *PrevFlag) SetBuilder(builder *strings.Builder) {
	(*p).builder = builder
}

func (p *PrevFlag) DoStaff(b bool, bytes []byte) {
	if !b {

	}
}

func (p *PrevFlag) Next(b bool) *IFlagChainNode {
	panic("implement me")
}

type Grep struct {
	reader io.Reader
	output strings.Builder
	regex  regexp.Regexp
	flag   IFlagChainNode
}

func (g *Grep) Match() error {
	bufReader := bufio.NewReader(g.reader)
	var (
		err     error
		slice   []byte
		matched bool
	)
	for err != io.EOF {
		slice, err = bufReader.ReadSlice('\n')
		if err != nil {
			return err
		}
		matched = g.regex.Match(slice)
		g.flag.DoStaff(matched, slice)
	}

	return nil
}

//TODO: responsive chain should consist task logic in DoStaff method,
// use single reader and delegate matching one to other
