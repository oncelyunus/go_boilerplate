package user

import (
	"context"

	"github.com/oncelyunus/go_boilerplate/pkg/internal"
)

const (
	_userCollection = "users"
)

type userMongoRepo struct {
	*internal.BaseMongoRepo
	collection string
}

func NewUserRepository(baseMongoRepo *internal.BaseMongoRepo) UserRepository {
	return &userMongoRepo{
		baseMongoRepo,
		_userCollection,
	}
}

// Create implements repository.UserRepository
func (umr *userMongoRepo) CreateUser(context.Context, *interface{}) (*interface{}, error) {
	panic("unimplemented")
}
