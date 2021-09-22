package pattern

import (
	"bufio"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Factory interface {
	Produce() Page
}

type PagePatternFactory struct {
	pattern string
}

func NewPagePatternFactory(pattern string) PagePatternFactory {
	return PagePatternFactory{pattern: pattern}
}

func (s PagePatternFactory) Produce() Page {
	page := new(StaticPage)
	page.pattern = s.pattern
	return page
}

type HtmlPageFactory struct {
	PathPrefix string
}

func (h HtmlPageFactory) Produce() Page {
	page := new(HtmlPage)
	page.PathPrefix = h.PathPrefix
	return page
}

func NewHtmlPageFactory(PathPrefix string) Factory {
	return HtmlPageFactory{PathPrefix: PathPrefix}
}

type Page interface {
	Init(*http.Request) error
	Render(http.ResponseWriter) error
}

type HtmlPage struct {
	PathPrefix string
	reader     *bufio.Reader
}

func (h *HtmlPage) Init(request *http.Request) error {
	name := strings.TrimPrefix(request.URL.Path, h.PathPrefix)

	file, err := os.OpenFile(name, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	(*h).reader = bufio.NewReader(file)
	return nil
}

func (h HtmlPage) Render(writer http.ResponseWriter) error {
	_, err := h.reader.WriteTo(writer)
	return err
}

type StaticPage struct {
	pattern string
	content []byte
}

func (s *StaticPage) Init(r *http.Request) error {
	vars := r.FormValue("id")

	if vars != "" {
		(*s).content = []byte(fmt.Sprintf(s.pattern, vars))
		return nil
	}
	return errors.New("wrong request")
}

func (s *StaticPage) Render(w http.ResponseWriter) error {
	_, err := w.Write(s.content)
	return err
}
