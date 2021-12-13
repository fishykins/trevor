package profile

type User struct {
	DiscordID string `json:"discordID" bson:"_id,omitempty"`
	SteamID   int64  `json:"steamID" bson:"steamID,omitempty"`
	Name      string `json:"name" bson:"name,omitempty"`
}
