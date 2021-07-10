package repo

import (
	"context"
	"example.com/app/database"
	"example.com/app/domain"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type StoryRepoImpl struct {
	story domain.Story
}

func (s StoryRepoImpl) Create(story domain.Story) error {
	conn := database.MongoConnectionPool.Get().(*database.Connection)
	defer database.MongoConnectionPool.Put(conn)

	_, err := conn.StoryCollection.InsertOne(context.TODO(), &story)

	if err != nil {
		return fmt.Errorf("error processing data")
	}

	return nil
}

func (s StoryRepoImpl) UpdateByID(story domain.Story) error {
	conn := database.MongoConnectionPool.Get().(*database.Connection)
	defer database.MongoConnectionPool.Put(conn)

	opts := options.FindOneAndUpdate().SetUpsert(true)
	filter := bson.D{{"_id", story.Id}}
	update := bson.D{{"$set", story}}

	conn.StoryCollection.FindOneAndUpdate(context.TODO(),
		filter, update, opts)

	return nil
}

func (s StoryRepoImpl) DeleteByID(story domain.Story) error {
	conn := database.MongoConnectionPool.Get().(*database.Connection)
	defer database.MongoConnectionPool.Put(conn)

	// sets mongo's read and write concerns
	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)

	// set up for a transaction
	session, err := conn.StartSession()

	if err != nil {
		panic(err)
	}

	defer session.EndSession(context.Background())

	// execute this code in a logical transaction
	callback := func(sessionContext mongo.SessionContext) (interface{}, error) {

		_, err = conn.StoryCollection.DeleteOne(context.TODO(), bson.D{{"_id", story.Id}})

		if err != nil {
			return nil, err

		}

		return nil, err
	}

	_, err = session.WithTransaction(context.Background(), callback, txnOpts)

	if err != nil {
		return fmt.Errorf("failed to delete user")
	}

	return nil
}

func NewStoryRepoImpl() StoryRepoImpl {
	var storyRepoImpl StoryRepoImpl

	return storyRepoImpl
}

