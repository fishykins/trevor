package core

import "fmt"

func UserTag(id string) string {
	return fmt.Sprintf("<@%s>", id)
}
