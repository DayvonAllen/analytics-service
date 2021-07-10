package repo

import "example.com/app/domain"

type StoryRepo interface {
	Create(story domain.Story) error
	UpdateByID(story domain.Story) error
	DeleteByID(story domain.Story) error
}
