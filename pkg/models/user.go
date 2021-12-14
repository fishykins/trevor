package models

type User struct {
	DiscordID uint64    `json:"ID" bson:"_id,omitempty"`
	SteamID   uint64    `json:"steamID" bson:"steamID,omitempty"`
	Name      string    `json:"name" bson:"name,omitempty"`
	Alignment Alignment `json:"alignment" bson:"alignment,omitempty"`
}
