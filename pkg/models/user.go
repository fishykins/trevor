package models

type User struct {
	ID        uint64    `json:"ID" bson:"_id,omitempty"`
	Name      string    `json:"name" bson:"name,omitempty"`
	Alignment Alignment `json:"alignment" bson:"alignment,omitempty"`
}
