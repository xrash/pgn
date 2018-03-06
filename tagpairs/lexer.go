package tagpairs

import (
	"fmt"
	"github.com/xrash/lexy"
)

func searchingOpenTagpair(l *lexy.Lexer, r rune) (lexy.State, error) {
	if lexy.IsEOF(r) {
		return nil, nil
	}

	if lexy.IsBlank(r) {
		return searchingOpenTagpair, nil
	}

	if r == '[' {
		return searchingTagpairKey, nil
	}

	return nil, fmt.Errorf("Expecting [, got %v", r)
}

func searchingTagpairKey(l *lexy.Lexer, r rune) (lexy.State, error) {
	if lexy.IsBlank(r) {
		return searchingTagpairKey, nil
	}

	if lexy.IsLetter(r) {
		l.Collect(r)
		return inTagpairKey, nil
	}

	return nil, fmt.Errorf("Expecting letter, got %v", r)
}

func inTagpairKey(l *lexy.Lexer, r rune) (lexy.State, error) {
	if lexy.IsBlank(r) {
		l.Emit("Key")
		return searchingTagpairValue, nil
	}

	if lexy.IsLetter(r) {
		l.Collect(r)
		return inTagpairKey, nil
	}

	return nil, fmt.Errorf("Expecting letter, got %v", r)
}

func searchingTagpairValue(l *lexy.Lexer, r rune) (lexy.State, error) {
	if lexy.IsBlank(r) {
		return searchingTagpairValue, nil
	}

	if r == '"' {
		return inTagpairValue, nil
	}

	return nil, fmt.Errorf(`Expecting ", got %v`, r)
}

func inTagpairValue(l *lexy.Lexer, r rune) (lexy.State, error) {
	if r == '"' {
		l.Emit("Value")
		return searchingCloseTagpair, nil
	}

	l.Collect(r)
	return inTagpairValue, nil
}

func searchingCloseTagpair(l *lexy.Lexer, r rune) (lexy.State, error) {
	if lexy.IsBlank(r) {
		return searchingCloseTagpair, nil
	}

	if r == ']' {
		return searchingOpenTagpair, nil
	}

	return nil, fmt.Errorf("Expecting ], got %v", r)
}
