package models

type User struct {
	DiscordID uint64            `json:"ID" bson:"_id,omitempty"`
	SteamID   uint64            `json:"steamID" bson:"steamID,omitempty"`
	Name      string            `json:"name" bson:"name,omitempty"`
	Alignment Alignment         `json:"alignment" bson:"alignment,omitempty"`
	Stats     map[string]int    `json:"stats" bson:"stats,omitempty"`
	Tokens    map[string]string `json:"tokens" bson:"tokens,omitempty"`
}

func NewUser(discordID uint64, steamID uint64, name string) *User {
	return &User{
		DiscordID: discordID,
		SteamID:   steamID,
		Name:      name,
		Alignment: Neutral(),
		Stats:     make(map[string]int),
		Tokens:    make(map[string]string),
	}
}

func (u *User) Init() {
	if u.Stats == nil {
		u.Stats = make(map[string]int)
	}
	if u.Tokens == nil {
		u.Tokens = make(map[string]string)
	}
}
