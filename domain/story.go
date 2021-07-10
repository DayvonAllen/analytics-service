package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Story struct {
	Id             primitive.ObjectID `bson:"_id" json:"id"`
	Title          string             `bson:"title" json:"title"`
	AuthorUsername string    		  `bson:"authorUsername" json:"authorUsername"`
}
