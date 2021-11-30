package lang

import "strings"

type Noun struct {
	Word string
}

func NewNoun(word string) Noun {
	return Noun{Word: strings.ToLower(word)}
}

func (n *Noun) String() string {
	return n.Word
}
