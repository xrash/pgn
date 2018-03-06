package pgn

import (
	"github.com/xrash/lexy"
	"io"
)

func searchingMovetextFirstLineBreak(l *lexy.Lexer, r rune) (lexy.State, error) {
	l.Collect(r)

	if lexy.IsLineBreak(r) {
		return searchingMovetextSecondLineBreak, nil
	}

	return searchingMovetextFirstLineBreak, nil
}

func searchingMovetextSecondLineBreak(l *lexy.Lexer, r rune) (lexy.State, error) {
	if lexy.IsLineBreak(r) {
		l.Emit("Movetext")
		return inTagpairs, nil
	}

	l.Collect(r)

	if lexy.IsBlank(r) {
		return searchingMovetextSecondLineBreak, nil
	}

	return searchingMovetextFirstLineBreak, nil
}

func inTagpairs(l *lexy.Lexer, r rune) (lexy.State, error) {
	l.Collect(r)

	if lexy.IsEOF(r) {
		l.Emit("Tagpairs")
		return inTagpairs, nil
	}

	return inTagpairs, nil
}

func SplitPGNSections(r io.Reader) (string, string, error) {
	tokens := make(chan *lexy.Token)

	var movetext, tagpairs string

	go func() {
		for t := range tokens {
			switch t.Key {
			case lexy.EOFToken:
				return
			case "Movetext":
				movetext = t.Value
			case "Tagpairs":
				tagpairs = t.Value
			}
		}
	}()

	l := lexy.NewLexer(tokens)
	if err := l.Do(r, searchingMovetextFirstLineBreak); err != nil {
		return "", "", err
	}

	return movetext, tagpairs, nil
}
