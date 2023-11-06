package auth

import (
	"github.com/gofiber/fiber/v2"
	"api-auth/db"
	"api-auth/api/models"
	"golang.org/x/crypto/bcrypt"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"api-auth/api/helpers"

)

func Register(c *fiber.Ctx) error {

	pool := db.Db().Database("api-auth").Collection("users")

	var newUser models.Auth

	c.BodyParser(&newUser)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"msg" : "Error hashing password",
		})
	}

	_, err = pool.InsertOne(context.TODO(), bson.D{
		{"email", newUser.Email},
		{"password", string(hashedPassword)},
		{"role", "user"},
		{"token", helpers.GenerateToken(newUser.Email)},
		{"created_at", time.Now()},
	})

	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"msg" : "Error creating user",
		})
	}

	return c.JSON(&fiber.Map{
		"msg": "User created",
	})

}

func Login ( c *fiber.Ctx ) error {

	pool := db.Db().Database("api-auth").Collection("users")

	var user models.Auth
	c.BodyParser(&user)
	var foundUser models.User
	
	err := pool.FindOne(context.TODO(), bson.D{
		{"email", user.Email},
	}).Decode(&foundUser)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(400).JSON(&fiber.Map{
				"msg": "User not found",
			})
		}

		return c.Status(400).JSON(&fiber.Map{
			"msg": "Error finding user",
		})
	}


	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))

	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"msg": "Incorrect password",
		})
	}



	token := helpers.GenerateJWT(foundUser.Id)


	return c.JSON(&fiber.Map{
		"msg": "Logged in",
		"token": token,
	})
}

func Profile(c *fiber.Ctx) error {

	user := c.Locals("user")

	return c.JSON(&fiber.Map{
		"user": user,
	})

}
