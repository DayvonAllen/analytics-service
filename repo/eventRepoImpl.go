package repo

import (
	"context"
	"example.com/app/database"
	"example.com/app/domain"
	"fmt"
)

type EventRepoImpl struct {
	event domain.Event
}

func (e EventRepoImpl) Create(event domain.Event) error {
	conn := database.MongoConnectionPool.Get().(*database.Connection)
	defer database.MongoConnectionPool.Put(conn)

	_, err := conn.EventCollection.InsertOne(context.TODO(), &event)

	if err != nil {
		return fmt.Errorf("error processing data")
	}

	return nil
}

func NewEventRepoImpl() EventRepoImpl {
	var eventRepoImpl EventRepoImpl

	return eventRepoImpl
}
