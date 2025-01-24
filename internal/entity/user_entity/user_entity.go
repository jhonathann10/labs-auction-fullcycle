package user_entity

import (
	"context"

	"github.com/jhonathann10/labs-auction-fullcycle/internal/internal_error"
)

type User struct {
	Id   string
	Name string
}

type UserRepositoryInterface interface {
	FindUserById(
		ctx context.Context, userId string) (*User, *internal_error.InternalError)
}
