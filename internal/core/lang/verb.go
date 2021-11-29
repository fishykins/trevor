package lang

type Verb struct {
	Word string
}

func NewVerb(inner string) Verb {
	return Verb{inner}
}

func (v *Verb) String() string {
	return v.Word
}
