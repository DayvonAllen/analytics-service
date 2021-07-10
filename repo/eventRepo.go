package repo

import "example.com/app/domain"

type EventRepo interface {
	Create(event domain.Event) error
}