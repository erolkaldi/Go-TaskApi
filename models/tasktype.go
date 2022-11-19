package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskType struct {
	ID         primitive.ObjectID `json:"id"`
	Code       string             `json:"code" validate:"required,min=2,max=20"`
	Name       string             `json:"name" validate:"required,min=2,max=100"`
	Deleted    bool               `json:"deleted"`
	RecordUser string             `json:"record_user"`
	RecordDate time.Time          `json:"record_date"`
}
