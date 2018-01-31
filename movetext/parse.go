package movetext

import (
	"github.com/xrash/lexy"
)

func Parse(content string) (*Movetext, error) {
	tokens := make(chan *lexy.Token, 1000)
	parserResult := make(chan *parserResult, 1)

	p := newParser(tokens)
	go p.parse(parserResult)

	l := lexy.NewLexer(tokens)
	err := l.DoString(content, searchingNumberOrHalfMoveOrResult)
	if err != nil {
		return nil, err
	}

	r := <-parserResult

	return r.movetext, r.err
}
