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
	Manage([]byte, bool)
}

type MatchedFlag struct {
	builder *strings.Builder
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
}

func (p *PostFlag) SetBuilder(builder *strings.Builder) {
	(*p).builder = builder
}

func (p *PostFlag) Manage(slice []byte, matched bool) {
	if matched { // match pattern, skip buffer writing
		(*p).matched = true
	} else if p.matched { // start writing in buffer until cap
		if len(p.buf) < cap(p.buf) {
			(*p).buf = append((*p).buf, slice)
		} else { // if len() == cap() stop writing in buffer
			(*p).matched = false
			for _, value := range p.buf {
				(*p).builder.Grow(len(slice))
				(*p).builder.Write(value)
			}
			(*p).buf = (*p).buf[:0]
		}
	}
}

func (p *PostFlag) Next() {

}

type CounterFlag struct {
	builder *strings.Builder
	next    IFlagChainNode
	index   uint
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
	buf     [][]byte
	builder *strings.Builder
	post    *PostFlag
	matched *MatchedFlag
}

func (p *PreFlag) Manage(slice []byte, matched bool) {
	if !matched {
		if len(p.buf) < cap(p.buf) {
			(*p).buf = append(p.buf, slice)
		} else {
			(*p).buf = (*p).buf[1:]
			(*p).buf = append(p.buf, slice)
		}
	} else {
		if p.post != nil {
			(*p.post).Manage(slice, matched)
		}
		for _, value := range p.buf {
			(*p).builder.Grow(len(value))
			(*p).builder.Write(value)
		}
		(*p).buf = (*p).buf[:0]
		(*p.matched).Manage(slice, matched)
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
}

func (p *Grep) SetRegex(regex *regexp.Regexp) {
	p.regex = regex
}

func (p *Grep) SetBuilder(builder *strings.Builder) {
	(*p).builder = builder
}

func (p *Grep) DoStaff(reader io.Reader) error {
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
			(*p).output.Manage(slice, matched)
		} else if err != io.EOF {
			fmt.Println(err)
		}
	}
	fmt.Println("\nOutuput:\n", (*p).builder.String())
	return err
}

//DONE: responsive chain should consist task logic in DoStaff method,
// use single reader and delegate matching one to other
