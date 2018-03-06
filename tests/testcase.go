package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/xrash/pgn"
	"github.com/xrash/pgn/movetext"
	"github.com/xrash/pgn/tagpairs"
	"testing"
)

type TestCase struct {
	input        string
	expectedGame *pgn.Game
}

func runTestCase(t *testing.T, c *TestCase) {
	g, err := pgn.ParseString(c.input)

	if !assert.Nil(t, err) {
		return
	}

	if assert.NotNil(t, g) {
		compareTagpairs(t, g.Tagpairs, c.expectedGame.Tagpairs)
		compareMovetext(t, g.Movetext, c.expectedGame.Movetext)
	}
}

func compareTagpairs(t *testing.T, v, e *tagpairs.Tagpairs) {
	assert.Equal(t, len(v.Data), len(e.Data), "Length of tagpairs data should be equal.")

	for key := range v.Data {
		msg := fmt.Sprintf("Tagpair for key %s not equal", key)
		assert.Equal(t, v.Data[key], e.Data[key], msg)
	}
}

func compareMovetext(t *testing.T, v, e *movetext.Movetext) {
	assert.Equal(t, len(e.Moves), len(v.Moves), "Length of movetext moves should be equal.")
	assert.Equal(t, e.Result, v.Result, "Movetext results should be equal.")

	for i := range v.Moves {
		assert.Equal(t, v.Moves[i].White, e.Moves[i].White, fmt.Sprintf("White move at index %d should be equal.", i))
		assert.Equal(t, v.Moves[i].Black, e.Moves[i].Black, fmt.Sprintf("Black move at index %d should be equal.", i))
		assert.Equal(t, v.Moves[i].Number, e.Moves[i].Number, fmt.Sprintf("Move number at index %s should be equal.", i))
	}
}
