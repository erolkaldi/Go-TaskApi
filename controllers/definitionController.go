package controllers

import (
	"context"
	"net/http"
	"task-api/database"
	"task-api/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var taskTypeCollection *mongo.Collection = database.GetCollection(database.Client, "tasktypes")
var validateTaskType = validator.New()

func AddTaskType() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		var taskType models.TaskType
		err := c.BindJSON(&taskType)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = validateTaskType.Struct(taskType)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "model is not valid"})
			return
		}
		count, err := taskTypeCollection.CountDocuments(ctx, bson.M{"code": taskType.Code})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Code is inuse"})
			return
		}
		taskType.ID = primitive.NewObjectID()
		taskType.Deleted = false
		taskType.RecordDate = time.Now().Add(time.Hour * 3)
		taskType.RecordUser = c.GetString("userid")

		ins, err := taskTypeCollection.InsertOne(ctx, taskType)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": true, "id": ins.InsertedID})

	}
}
