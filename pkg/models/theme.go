package models

type Theme struct {
	ID          int32       `json:"ID" bson:"_id,omitempty"`
	Name        string      `json:"name" bson:"name,omitempty"`
	Description string      `json:"description" bson:"description,omitempty"`
	Author      string      `json:"author" bson:"author,omitempty"`
	Links       []ThemeLink `json:"links" bson:"links,omitempty"`
}

type ThemeLink struct {
	A        int32 `json:"a" bson:"a,omitempty"`
	B        int32 `json:"b" bson:"b,omitempty"`
	Strength int32 `json:"strength" bson:"strength,omitempty"`
}
