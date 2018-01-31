package pgn

import (
	"bufio"
	"github.com/xrash/pgn/movetext"
	"github.com/xrash/pgn/tagpairs"
	"io"
	"strings"
)

func ParseString(pgn string) (*Game, error) {
	return Parse(strings.NewReader(strings.TrimSpace(pgn)))
}

func Parse(r io.Reader) (*Game, error) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	tagpairSection := scanUntilBlankLine(s)
	movetextSection := scanUntilBlankLine(s)

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

func scanUntilBlankLine(s *bufio.Scanner) string {
	lines := make([]string, 0)

	for s.Scan() {
		line := s.Text()

		if line == "" {
			break
		}

		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}
