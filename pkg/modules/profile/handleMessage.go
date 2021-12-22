package profile

import (
	"strings"

	"github.com/fishykins/trevor/pkg/models"
)

func HandleMessage(user *models.User, msg string) {
	user.Stats["wordCount"] += len(strings.Fields(msg))
	user.Stats["messageCount"]++
	UpdateUser(user)
}
