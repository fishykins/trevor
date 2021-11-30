package lang

import "strings"

type Verb struct {
	Word string
}

func NewVerb(inner string) Verb {
	return Verb{Word: strings.ToLower(inner)}
}

func (v *Verb) String() string {
	return v.Word
}
