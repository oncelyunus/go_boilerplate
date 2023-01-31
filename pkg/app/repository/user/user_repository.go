package user

import "context"

type UserRepository interface {
	CreateUser(context.Context, *interface{}) (*interface{}, error)
}
