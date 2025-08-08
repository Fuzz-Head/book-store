package authutil

func ScopesForRole(role string) []string {
	switch role {
	case "admin":
		return []string{
			"can:read:book",
			"can:read:books",
			"can:create:book",
			"can:update:book",
			"can:delete:book",
			"can:read:author",
			"can:read:authors",
			"can:create:author",
			"can:update:author",
			"can:delete:author",
		}
	case "superUser":
		return []string{
			"can:read:book",
			"can:read:books",
			"can:create:book",
			"can:update:book",
		}
	case "user":
		return []string{
			"can:read:book",
			"can:read:books",
			"can:read:author",
			"can:read:authors",
		}
	case "author":
		return []string{
			"can:read:book",
			"can:read:books",
			"can:read:author",
			"can:read:authors",
			"can:update:author",
		}
	default:
		return []string{}
	}
}
