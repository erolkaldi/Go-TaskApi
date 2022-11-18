package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Email    *string            `json:"email" validate:"required"`
	UserName *string            `json:"user_name" validate:"required,min=2,max=255"`
	Password *string            `json:"password" validate:"required"`
	UserType *string            `json:"user_type" validate:"required,eq=ADMIN|eq=GUEST|eq=USER"`
	UserID   string             `json:"user_id"`
	Deleted  bool               `json:"deleted"`
}
