package pgn

import (
	"github.com/xrash/pgn/movetext"
	"github.com/xrash/pgn/tagpairs"
)

type Game struct {
	Tagpairs *tagpairs.Tagpairs
	Movetext *movetext.Movetext
}
