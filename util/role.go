package util

import (
	"panel/ent"
)

func IsRoleDeleted(roleData *ent.Role) bool {
	return roleData.DeletedAt.Unix() != -62135596800
}
