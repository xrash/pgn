package movetext

import (
	"fmt"
	"github.com/xrash/lexy"
)

func searchingNumberOrHalfMoveOrResult(l *lexy.Lexer, r rune) (lexy.State, error) {
	if lexy.IsEOF(r) {
		return searchingNumberOrHalfMoveOrResult, nil
	}

	if lexy.IsBlank(r) {
		return searchingNumberOrHalfMoveOrResult, nil
	}

	if r == '{' {
		return inBraceCommentary, nil
	}

	if r == '(' {
		return inParenthesisCommentary, nil
	}

	if r == '*' {
		l.Collect(r)
		l.Emit("Result")
		return searchingNumberOrHalfMoveOrResult, nil
	}

	if lexy.IsNumber(r) {
		l.Collect(r)
		return inNumberOrResult, nil
	}

	if lexy.IsLetter(r) {
		l.Collect(r)
		return inHalfMove, nil
	}

	return nil, fmt.Errorf("Expecting a number, a command or a result, got %v", r)
}

func inBraceCommentary(l *lexy.Lexer, r rune) (lexy.State, error) {
	if r == '}' {
		return searchingNumberOrHalfMoveOrResult, nil
	}

	return inBraceCommentary, nil
}

func inParenthesisCommentary(l *lexy.Lexer, r rune) (lexy.State, error) {
	if r == ')' {
		return searchingNumberOrHalfMoveOrResult, nil
	}

	return inParenthesisCommentary, nil
}

func inNumberOrResult(l *lexy.Lexer, r rune) (lexy.State, error) {
	if lexy.IsBlank(r) {
		l.Emit("MoveNumber")
		return searchingNumberOrHalfMoveOrResult, nil
	}

	if r == '/' || r == '-' {
		l.Collect(r)
		return inResult, nil
	}

	if lexy.IsNumber(r) {
		l.Collect(r)
		return inNumberOrResult, nil
	}

	if r == '.' {
		l.Collect(r)
		return inDots, nil
	}

	return nil, fmt.Errorf("Expecting a number or a result, got %v", r)
}

func inDots(l *lexy.Lexer, r rune) (lexy.State, error) {
	if r == '.' {
		l.Collect(r)
		return inDots, nil
	}

	if lexy.IsBlank(r) {
		l.Emit("MoveNumber")
		return searchingNumberOrHalfMoveOrResult, nil
	}

	l.Emit("MoveNumber")
	l.Collect(r)
	return inHalfMove, nil
}

func inResult(l *lexy.Lexer, r rune) (lexy.State, error) {
	if lexy.IsBlank(r) || lexy.IsEOF(r) {
		l.Emit("Result")
		return searchingNumberOrHalfMoveOrResult, nil
	}

	if lexy.IsNumber(r) || r == '-' || r == '/' {
		l.Collect(r)
		return inResult, nil
	}

	return nil, fmt.Errorf("Expecting a result, got %v", r)
}

func inHalfMove(l *lexy.Lexer, r rune) (lexy.State, error) {
	if lexy.IsBlank(r) {
		l.Emit("HalfMove")
		return searchingNumberOrHalfMoveOrResult, nil
	}

	l.Collect(r)

	return inHalfMove, nil
}
