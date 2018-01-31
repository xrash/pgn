package tagpairs

type Tagpairs struct {
	Data map[string]string
}

func NewTagpairs() *Tagpairs {
	return &Tagpairs{
		Data: make(map[string]string),
	}
}
