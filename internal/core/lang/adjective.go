package lang

type Adjective struct {
	Word string
}

func NewAdjective(word string) Adjective {
	return Adjective{Word: word}
}

func (a *Adjective) String() string {
	return a.Word
}
