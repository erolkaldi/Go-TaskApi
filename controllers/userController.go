package controllers

import (
	"context"
	"net/http"
	"task-api/database"
	"task-api/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection = database.GetCollection(database.Client, "users")

func GetUser() gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		id := c.Param("id")
		var user models.User
		err := collection.FindOne(ctx, bson.M{"userid": id}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		user.Password = nil
		c.JSON(http.StatusOK, user)
	}

}
func GetUsers() gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		var users []models.User

		cursor, err := collection.Find(ctx, bson.D{})
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		for cursor.Next(ctx) {
			var user models.User
			err := cursor.Decode(&user)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			user.Password = nil
			users = append(users, user)

		}
		c.JSON(http.StatusOK, users)
	}

}
