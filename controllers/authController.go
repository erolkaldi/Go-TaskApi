package controllers

import (
	"context"
	"net/http"
	"task-api/database"
	"time"

	"task-api/models"

	"task-api/helpers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.GetCollection(database.Client, "users")
var validate = validator.New()

func RegisterUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var register models.RegisterDto
		if err := c.BindJSON(&register); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Model not valid"})
			return
		}
		validationErr := validate.Struct(register)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": register.Email})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
			return
		}
		hasHashPassword := helpers.HashPassword(register.Password)
		var user models.User
		user.Email = &register.Email
		user.UserName = &register.Username
		user.Password = &hasHashPassword
		user.UserType = &register.UserType
		user.ID = primitive.NewObjectID()
		user.UserID = user.ID.Hex()
		user.Deleted = false
		insertResult, err := userCollection.InsertOne(ctx, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, insertResult)
	}
}

func GetToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var loginDto models.Login
		var user models.User
		if err := c.BindJSON(&loginDto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Model not valid"})
			return
		}
		err := userCollection.FindOne(ctx, bson.M{"email": loginDto.Email}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		valid, err := helpers.ComparePassword(loginDto.Password, *user.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if !valid {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
			return
		}
		var token models.Token
		token.UserID = user.UserID
		token.UserName = *user.UserName
		token.UserType = *user.UserType
		token.Email = *user.Email
		acess_token, refresh_token, err := helpers.GenerateTokens(token.Email, token.UserName, token.UserType, token.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		token.AccessToken = acess_token
		token.RefreshToken = refresh_token
		c.JSON(http.StatusOK, token)

	}

}
