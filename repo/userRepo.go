package repo

import (
	"example.com/app/domain"
)

type UserRepo interface {
	Create(domain.User) error
	UpdateByID(domain.User) error
	DeleteByID(domain.User) error
}
