package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id                   primitive.ObjectID `bson:"_id" json:"id"`
	Username             string             `bson:"username" json:"username"`
}
