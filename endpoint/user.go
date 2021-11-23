package endpoint

import (
	"dota2updates-service/commom"
	"dota2updates-service/util"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var users []commom.User
var secretKey string

func (a *Api) SignUp(c *gin.Context) {
	var signUp commom.User
	err := c.ShouldBind(&signUp)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal error"})
		return
	}

	token := generateToken(signUp.Email)
	signUp.Token = token

	users = append(users, signUp)
}

func getSecretKey() string {
	if secretKey == "" {
		secretKey = util.RandomString(15)
	}

	return secretKey
}

func generateToken(email string) string {
	claims := commom.AuthClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(getSecretKey()))
	if err != nil {
		panic(err)
	}
	return t
}

func validateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(getSecretKey()), nil
	})
}
