package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type IFlagChainNode interface {
	SetBuilder(*strings.Builder)
	SetNext(IFlagChainNode)
	Manage([]byte, bool)
}

type MatchedFlag struct {
	builder *strings.Builder
	next    IFlagChainNode
}

func (m *MatchedFlag) SetNext(next IFlagChainNode) {
	m.next = next
}

func (m *MatchedFlag) SetBuilder(builder *strings.Builder) {
	(*m).builder = builder
}

func (m *MatchedFlag) Manage(slice []byte, matched bool) {
	if matched {
		(*m).builder.Grow(len(slice))
		(*m).builder.Write(slice)
	}
}

type PostFlag struct {
	builder *strings.Builder
	buf     [][]byte
	matched bool
	next    IFlagChainNode
}

func (p *PostFlag) Matched() bool {
	return p.matched
}

func (p *PostFlag) SetBuilder(builder *strings.Builder) {
	(*p).builder = builder
}

func (p *PostFlag) Manage(slice []byte, matched bool) {
	switch p.next.(type) {
	case *PreFlag:
		p.next.(*PreFlag).SetMatchedPost(p.Matched())
	}

	//fmt.Printf("POSTFLAG: input match %v, own match %v, next slice %s\n", next,p.next, string(slice))
	if matched { // match pattern, skip buffer writing
		(*p).matched = true
	} else if p.matched { // start writing in buffer until cap
		if len(p.buf) < cap(p.buf) {
			//fmt.Println("appended")
			(*p).buf = append((*p).buf, slice)
			//fmt.Println(len((*p).buf))
			(*p).builder.Grow(len(slice))
			(*p).builder.Write(slice)
			if len(p.buf) == cap(p.buf) {
				(*p).matched = false
				(*p).buf = (*p).buf[:0]
			}
		}
	}
	if p.next != nil {
		p.next.Manage(slice, matched)
	}

}

func (p *PostFlag) SetNext(node IFlagChainNode) {
	(*p).next = node
}

type CounterFlag struct {
	builder *strings.Builder
	next    IFlagChainNode
	index   uint
}

func (c *CounterFlag) SetNext(next IFlagChainNode) {
	c.next = next
}

func (c *CounterFlag) Index() uint {
	return c.index
}

func (c *CounterFlag) SetBuilder(builder *strings.Builder) {
	(*c).builder = builder
}

func (c *CounterFlag) Manage(slice []byte, matched bool) {
	(*c).index++
	slice = append([]byte(strconv.Itoa(int(c.index))+":"), slice...)
	(*c).next.Manage(slice, matched)
}

type PreFlag struct {
	buf         [][]byte
	builder     *strings.Builder
	next        IFlagChainNode
	matchedPost bool
}

func (p *PreFlag) SetNext(next IFlagChainNode) {
	p.next = next
}

func (p *PreFlag) SetMatchedPost(matchedPost bool) {
	p.matchedPost = matchedPost
}

func (p *PreFlag) Manage(slice []byte, matched bool) {

	if !matched && !p.matchedPost {
		if len(p.buf) < cap(p.buf) {
			(*p).buf = append(p.buf, slice)
		} else {
			(*p).buf = (*p).buf[1:]
			(*p).buf = append(p.buf, slice)
		}
	} else {
		for _, value := range p.buf {
			(*p).builder.Grow(len(value))
			(*p).builder.Write(value)
		}
		(*p).buf = (*p).buf[:0]
		(*p).next.Manage(slice, matched)
		//(*p).builder.Write(slice)
	}
}

func (p *PreFlag) SetBuilder(builder *strings.Builder) {
	(*p).builder = builder
}

type Grep struct {
	builder *strings.Builder
	regex   *regexp.Regexp
	output  IFlagChainNode
	exclude bool
}

func (p *Grep) SetRegex(regex *regexp.Regexp) {
	p.regex = regex
}

func (p *Grep) SetBuilder(builder *strings.Builder) {
	(*p).builder = builder
}

func (p *Grep) DoStaff(reader io.Reader) (error, string) {
	bufReader := bufio.NewReader(reader)
	var (
		err     error
		slice   []byte
		matched bool
	)
	for err != io.EOF {
		slice, err = bufReader.ReadSlice('\n')
		if err == nil {
			matched = (*p).regex.Match(slice)
			if p.exclude {
				matched = !matched
			}
			(*p).output.Manage(slice, matched)
		} else if err != io.EOF {
			fmt.Println(err)
		}
	}
	return err, (*p).builder.String()
}
