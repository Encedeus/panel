package util

import (
	"panel/ent"
)

func IsUserDeleted(userData *ent.User) bool {
	return userData.DeletedAt.Unix() != -62135596800
}
