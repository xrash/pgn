package tagpairs

import (
	"github.com/xrash/lexy"
)

func Parse(content string) (*Tagpairs, error) {
	tokens := make(chan *lexy.Token, 1000)
	parserResult := make(chan *parserResult, 1)

	p := newParser(tokens)
	go p.parse(parserResult)

	l := lexy.NewLexer(tokens)
	err := l.DoString(content, searchingOpenTagpair)
	if err != nil {
		return nil, err
	}

	r := <-parserResult

	return r.tagpairs, r.err
}
