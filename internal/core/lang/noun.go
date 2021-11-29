package lang

type Noun struct {
	Word string
}

func NewNoun(word string) Noun {
	return Noun{Word: word}
}

func (n *Noun) String() string {
	return n.Word
}
