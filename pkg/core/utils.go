package core

import "fmt"

func UserTag(id string) string {
	return fmt.Sprintf("<@%s>", id)
}

func BangTag(id string) string {
	return fmt.Sprintf("<@!%s>", id)
}
