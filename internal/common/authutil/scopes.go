package authutil

func ScopesForRole(role string) []string {
	switch role {
	case "admin":
		return []string{
			"can:read:books",
			"can:read:book",
			"can:create:book",
			"can:update:book",
			"can:delete:book",
		}
	case "superUser":
		return []string{
			"can:read:books",
			"can:read:book",
			"can:create:book",
			"can:update:book",
		}
	case "user":
		return []string{
			"can:read:book",
			"can:read:books",
		}
	default:
		return []string{}
	}
}
