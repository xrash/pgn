package pgn

import (
	"github.com/xrash/pgn/movetext"
	"github.com/xrash/pgn/tagpairs"
	"io"
	"strings"
)

func ParseString(pgn string) (*Game, error) {
	return Parse(strings.NewReader(pgn))
}

func Parse(r io.Reader) (*Game, error) {
	tagpairSection, movetextSection, err := SplitPGNSections(r)
	if err != nil {
		return nil, err
	}

	tagpairs, err := tagpairs.Parse(tagpairSection)
	if err != nil {
		return nil, err
	}

	movetext, err := movetext.Parse(movetextSection)
	if err != nil {
		return nil, err
	}

	g := &Game{
		Tagpairs: tagpairs,
		Movetext: movetext,
	}

	return g, nil
}
