package models

type UserClaims struct {
	Role   string   `json:"role"`
	Scopes []string `json:"scopes"`
}
