package movetext

import (
	"fmt"
	"github.com/xrash/lexy"
	"strconv"
	"strings"
)

type parserState func(*lexy.Token, *Movetext) (parserState, error)

type parserResult struct {
	movetext *Movetext
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
	var state parserState = expectingWhiteMoveNumber
	var err error
	movetext := NewMovetext()

	for t := range p.tokens {
		if t.Key == lexy.EOFToken {
			break
		}

		if t.Key == lexy.ErrorToken {
			err = fmt.Errorf("Error parsing %s: %v", t.Value, err)
			break
		}

		state, err = state(t, movetext)
		if err != nil {
			err = fmt.Errorf("Error parsing %s: %v", t.Value, err)
			break
		}
	}

	back <- &parserResult{
		movetext: movetext,
		err:      err,
	}
}

func expectingWhiteMoveNumber(t *lexy.Token, m *Movetext) (parserState, error) {
	if t.Key == "Result" {
		m.Result = Result(t.Value)
		return expectingNothing, nil
	}

	if t.Key != "MoveNumber" {
		return nil, fmt.Errorf("Expecting white move number, got %v", t)
	}

	v := strings.TrimRight(t.Value, ".")
	moveNumber, err := strconv.Atoi(v)
	if err != nil {
		return nil, fmt.Errorf("Could not understand move number %s", t.Value)
	}

	move := &Move{
		Number: moveNumber,
	}

	m.Moves = append(m.Moves, move)

	return expectingWhiteHalfMove, nil
}

func expectingWhiteHalfMove(t *lexy.Token, m *Movetext) (parserState, error) {
	if t.Key != "HalfMove" {
		return nil, fmt.Errorf("Expecting white half move, got %v", t)
	}

	m.Moves[len(m.Moves)-1].White = t.Value

	return expectingBlackMoveNumberOrHalfMove, nil
}

func expectingBlackMoveNumberOrHalfMove(t *lexy.Token, m *Movetext) (parserState, error) {
	if t.Key == "Result" {
		m.Result = Result(t.Value)
		return expectingNothing, nil
	}

	if t.Key != "MoveNumber" && t.Key != "HalfMove" {
		return nil, fmt.Errorf("Expecting black move number or half move, got %v", t)
	}

	if t.Key == "MoveNumber" {
		return expectingBlackHalfMove, nil
	}

	// HalfMove
	m.Moves[len(m.Moves)-1].Black = t.Value
	return expectingWhiteMoveNumber, nil
}

func expectingBlackHalfMove(t *lexy.Token, m *Movetext) (parserState, error) {
	if t.Key != "HalfMove" {
		return nil, fmt.Errorf("Expecting black half move, got %v", t)
	}

	m.Moves[len(m.Moves)-1].Black = t.Value

	return expectingWhiteMoveNumber, nil
}

func expectingNothing(t *lexy.Token, m *Movetext) (parserState, error) {
	msg := fmt.Sprintf("Found %v after result.", t)
	return nil, fmt.Errorf(msg)
}
