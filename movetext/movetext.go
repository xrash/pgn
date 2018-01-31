package movetext

type Result string

const (
	WhiteWins Result = "1-0"
	BlackWins Result = "0-1"
	Draw      Result = "1/2-1/2"
	Unknown   Result = "*"
)

type Movetext struct {
	Moves  []*Move
	Result Result
}

func NewMovetext() *Movetext {
	return &Movetext{
		Moves: make([]*Move, 0),
	}
}

type Move struct {
	Number int
	White  string
	Black  string
}
