package lang

import "strings"

type Adjective struct {
	Word string
}

func NewAdjective(word string) Adjective {
	return Adjective{Word: strings.ToLower(word)}
}

func (a *Adjective) String() string {
	return a.Word
}
