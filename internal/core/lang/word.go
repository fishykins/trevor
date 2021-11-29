package lang

type Word struct {
	ID          int32  `json:"ID" bson:"_id,omitempty"`
	Word        string `json:"word" bson:"word,omitempty"`
	WordType    string `json:"type" bson:"type,omitempty"`
	Source      string `json:"source" bson:"source,omitempty"`
	Description string `json:"description" bson:"description,omitempty"`
	Themes      []int  `json:"themes" bson:"themes,omitempty"`
}
