package tests

import (
	"github.com/xrash/pgn"
	"github.com/xrash/pgn/movetext"
	"github.com/xrash/pgn/tagpairs"
)

var __testcase_1 *TestCase

func init() {
	createTestCase1()
}

func createTestCase1() {
	tagpairs := &tagpairs.Tagpairs{
		Data: map[string]string{
			"Event":  "F/S Return Match",
			"Site":   "Belgrade, Serbia JUG",
			"Date":   "1992.11.04",
			"Round":  "29",
			"White":  "Fischer, Robert J.",
			"Black":  "Spassky, Boris V.",
			"Result": "1/2-1/2",
		},
	}

	moves := []*movetext.Move{
		&movetext.Move{
			Number: 1,
			White:  "e4",
			Black:  "e5",
		},
		&movetext.Move{
			Number: 2,
			White:  "Nf3",
			Black:  "Nc6",
		},
		&movetext.Move{
			Number: 3,
			White:  "Bb5",
			Black:  "a6",
		},
		&movetext.Move{
			Number: 4,
			White:  "Ba4",
			Black:  "Nf6",
		},
		&movetext.Move{
			Number: 5,
			White:  "O-O",
			Black:  "Be7",
		},
		&movetext.Move{
			Number: 6,
			White:  "Re1",
			Black:  "b5",
		},
		&movetext.Move{
			Number: 7,
			White:  "Bb3",
			Black:  "d6",
		},
		&movetext.Move{
			Number: 8,
			White:  "c3",
			Black:  "O-O",
		},
		&movetext.Move{
			Number: 9,
			White:  "h3",
			Black:  "Nb8",
		},
		&movetext.Move{
			Number: 10,
			White:  "d4",
			Black:  "Nbd7",
		},
		&movetext.Move{
			Number: 11,
			White:  "c4",
			Black:  "c6",
		},
		&movetext.Move{
			Number: 12,
			White:  "cxb5",
			Black:  "axb5",
		},
		&movetext.Move{
			Number: 13,
			White:  "Nc3",
			Black:  "Bb7",
		},
		&movetext.Move{
			Number: 14,
			White:  "Bg5",
			Black:  "b4",
		},
		&movetext.Move{
			Number: 15,
			White:  "Nb1",
			Black:  "h6",
		},
		&movetext.Move{
			Number: 16,
			White:  "Bh4",
			Black:  "c5",
		},
		&movetext.Move{
			Number: 17,
			White:  "dxe5",
			Black:  "Nxe4",
		},
		&movetext.Move{
			Number: 18,
			White:  "Bxe7",
			Black:  "Qxe7",
		},
		&movetext.Move{
			Number: 19,
			White:  "exd6",
			Black:  "Qf6",
		},
		&movetext.Move{
			Number: 20,
			White:  "Nbd2",
			Black:  "Nxd6",
		},
		&movetext.Move{
			Number: 21,
			White:  "Nc4",
			Black:  "Nxc4",
		},
		&movetext.Move{
			Number: 22,
			White:  "Bxc4",
			Black:  "Nb6",
		},
		&movetext.Move{
			Number: 23,
			White:  "Ne5",
			Black:  "Rae8",
		},
		&movetext.Move{
			Number: 24,
			White:  "Bxf7+",
			Black:  "Rxf7",
		},
		&movetext.Move{
			Number: 25,
			White:  "Nxf7",
			Black:  "Rxe1+",
		},
		&movetext.Move{
			Number: 26,
			White:  "Qxe1",
			Black:  "Kxf7",
		},
		&movetext.Move{
			Number: 27,
			White:  "Qe3",
			Black:  "Qg5",
		},
		&movetext.Move{
			Number: 28,
			White:  "Qxg5",
			Black:  "hxg5",
		},
		&movetext.Move{
			Number: 29,
			White:  "b3",
			Black:  "Ke6",
		},
		&movetext.Move{
			Number: 30,
			White:  "a3",
			Black:  "Kd6",
		},
		&movetext.Move{
			Number: 31,
			White:  "axb4",
			Black:  "cxb4",
		},
		&movetext.Move{
			Number: 32,
			White:  "Ra5",
			Black:  "Nd5",
		},
		&movetext.Move{
			Number: 33,
			White:  "f3",
			Black:  "Bc8",
		},
		&movetext.Move{
			Number: 34,
			White:  "Kf2",
			Black:  "Bf5",
		},
		&movetext.Move{
			Number: 35,
			White:  "Ra7",
			Black:  "g6",
		},
		&movetext.Move{
			Number: 36,
			White:  "Ra6+",
			Black:  "Kc5",
		},
		&movetext.Move{
			Number: 37,
			White:  "Ke1",
			Black:  "Nf4",
		},
		&movetext.Move{
			Number: 38,
			White:  "g3",
			Black:  "Nxh3",
		},
		&movetext.Move{
			Number: 39,
			White:  "Kd2",
			Black:  "Kb5",
		},
		&movetext.Move{
			Number: 40,
			White:  "Rd6",
			Black:  "Kc5",
		},
		&movetext.Move{
			Number: 41,
			White:  "Ra6",
			Black:  "Nf2",
		},
		&movetext.Move{
			Number: 42,
			White:  "g4",
			Black:  "Bd3",
		},
		&movetext.Move{
			Number: 43,
			White:  "Re6",
			Black:  "",
		},
	}

	movetext := &movetext.Movetext{
		Moves:  moves,
		Result: "1/2-1/2",
	}

	game := &pgn.Game{
		Tagpairs: tagpairs,
		Movetext: movetext,
	}

	__testcase_1 = &TestCase{
		input:        __fischer_vs_spassky_game_29,
		expectedGame: game,
	}
}
