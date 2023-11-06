package helpers

import (
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/golang-jwt/jwt"
	"time"
)

func GenerateToken (email string) string {
	
	hashedEmail, err := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(hashedEmail)

}

func GenerateJWT (id primitive.ObjectID) string {
	
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)


	claims["id"] = id.Hex()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		panic(err)
	}

	return tokenString


}