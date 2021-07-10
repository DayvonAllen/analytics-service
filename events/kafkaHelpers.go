package events

import (
	"example.com/app/domain"
	"example.com/app/repo"
	"fmt"
)

func ProcessMessage(message domain.Message) error {

	if message.ResourceType == "user" {
		// 201 is the created messageType
		if message.MessageType == 201 {
			user := message.User
			err := repo.UserRepoImpl{}.Create(user)

			if err != nil {
				return err
			}
			return nil
		}

		// 200 is the updated messageType
		if message.MessageType == 200 {
			user := message.User

			err := repo.UserRepoImpl{}.UpdateByID(user)
			if err != nil {
				return err
			}
			return nil
		}

		// 204 is the deleted messageType
		if message.MessageType == 204 {
			user := message.User

			err := repo.UserRepoImpl{}.DeleteByID(user)

			if err != nil {
				return err
			}
			return nil
		}
	}

	if message.ResourceType == "story" {
		// 201 is the created messageType
		if message.MessageType == 201 {
			story := message.Story
			err := repo.StoryRepoImpl{}.Create(story)
			if err != nil {
				return err
			}
			return nil
		}

		// 200 is the updated messageType
		if message.MessageType == 200 {
			story := message.Story

			err := repo.StoryRepoImpl{}.UpdateByID(story)
			if err != nil {
				return err
			}
			return nil
		}

		// 204 is the deleted messageType
		if message.MessageType == 204 {
			story := message.Story

			err := repo.StoryRepoImpl{}.DeleteByID(story)

			if err != nil {
				return err
			}
			return nil
		}
	}

	if message.ResourceType == "event" {
		event := message.Event
		err := repo.EventRepoImpl{}.Create(event)
		if err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("cannot process this message")

}
