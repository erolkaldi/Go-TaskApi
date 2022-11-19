package controllers

import (
	"context"
	"net/http"
	"task-api/database"
	"task-api/helpers"
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

func ChangePassword() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		id := c.GetString("userid")

		var user models.User
		err := collection.FindOne(ctx, bson.M{"userid": id}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		var changeDto models.ChangePassword
		err = c.BindJSON(&changeDto)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		valid, err := helpers.ComparePassword(changeDto.OldPassword, *user.Password)
		if err != nil || !valid {
			c.JSON(http.StatusBadRequest, gin.H{"error": "old password not valid"})
			return
		}
		hashed := helpers.HashPassword(changeDto.NewPassword)
		filter := bson.M{"userid": bson.M{"$eq": id}}
		update := bson.M{"$set": bson.M{"password": &hashed}}
		result, err := collection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if result.ModifiedCount == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database update error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": true})

	}
}
