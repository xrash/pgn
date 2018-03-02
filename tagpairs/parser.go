package tagpairs

import (
	"fmt"
	"github.com/xrash/lexy"
)

var __lastKey string

type parserState func(*lexy.Token, *Tagpairs) (parserState, error)

type parserResult struct {
	tagpairs *Tagpairs
	err      error
}

type parser struct {
	tokens chan *lexy.Token
}

func newParser(tokens chan *lexy.Token) *parser {
	return &parser{
		tokens: tokens,
	}
}

func (p *parser) parse(back chan *parserResult) {
	var state parserState = expectingKey
	var err error
	tagpairs := NewTagpairs()

	for t := range p.tokens {
		if t.Key == lexy.ErrorToken {
			back <- &parserResult{
				err: fmt.Errorf("%s", t.Value),
			}
			return
		}

		if t.Key == lexy.EOFToken {
			back <- &parserResult{
				tagpairs: tagpairs,
			}
			return
		}

		state, err = state(t, tagpairs)
		if err != nil {
			back <- &parserResult{
				err: fmt.Errorf("%s", t.Value),
			}
			return
		}
	}
}

func expectingKey(t *lexy.Token, tp *Tagpairs) (parserState, error) {
	if t.Key == "Key" {
		__lastKey = t.Value
		return expectingValue, nil
	}

	return nil, fmt.Errorf("Expected key, got: %v", t)
}

func expectingValue(t *lexy.Token, tp *Tagpairs) (parserState, error) {
	if t.Key == "Value" {
		tp.Data[__lastKey] = t.Value
		return expectingKey, nil
	}

	return nil, fmt.Errorf("Expected value, got: %v", t)
}
