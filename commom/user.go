package commom

import "github.com/golang-jwt/jwt"

type User struct {
	Nickname string `json:"nickname"`
	SteamID  string `json:"steam_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string
}

type AuthClaims struct {
	Email string
	jwt.StandardClaims
}
