package lang

import "github.com/fishykins/trevor/pkg/models"

type Word struct {
	ID          int32            `json:"ID" bson:"_id,omitempty"`
	Inner       string           `json:"word" bson:"word,omitempty"`
	Type        string           `json:"type" bson:"type,omitempty"`
	Source      string           `json:"source" bson:"source,omitempty"`
	Description string           `json:"description" bson:"description,omitempty"`
	Themes      []int            `json:"themes" bson:"themes,omitempty"`
	Alignment   models.Alignment `json:"alignment" bson:"alignment,omitempty"`
}

func (w *Word) IntoAdjective() Adjective {
	if w.Type != "adjective" {
		panic("Word is not an adjective: " + w.Type)
	}
	return NewAdjective(w.Inner)
}

func (w *Word) IntoNoun() Noun {
	if w.Type != "noun" {
		panic("Word is not a noun: " + w.Type)
	}
	return NewNoun(w.Inner)
}

func (w *Word) IntoVerb() Verb {
	if w.Type != "verb" {
		panic("Word is not a verb: " + w.Type)
	}
	return NewVerb(w.Inner)
}
