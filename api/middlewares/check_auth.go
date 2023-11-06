package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"context"
	"strings"
	"api-auth/db"
	"go.mongodb.org/mongo-driver/bson"
	"api-auth/api/models"
)

func CheckAuth(c *fiber.Ctx) error {

	// Primero, verificas si el encabezado de autorización está presente
	authHeader := c.Get("Authorization")
	if authHeader == ""{
		return c.Status(401).JSON(&fiber.Map{
			"msg": "No token provided",
		})
	}

	// divide el encabezado de autorizacion en dos partes, "Bearer" y el token

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(401).JSON(&fiber.Map{
			"msg": "Invalid token",
		})
	}

	token := parts[1]


	// Ahora, verificas si el token es válido
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return c.Status(401).JSON(&fiber.Map{
			"msg": "Invalid token",
		})
	}


	// Finalmente, verificas si el usuario existe en la base de datos

	pool := db.Db().Database("api-auth").Collection("users")

	userID, err := primitive.ObjectIDFromHex(claims["id"].(string))
	if err != nil {
		return c.Status(401).JSON(&fiber.Map{
			"msg": "Invalid token",
		})
	}


	var user models.User


	err = pool.FindOne(context.TODO(), bson.D{{"_id", userID}}).Decode(&user)



	if err != nil {
		return c.Status(401).JSON(&fiber.Map{
			"msg": "user not found",
		})
	}

	c.Locals("user", user)



	return c.Next()
}