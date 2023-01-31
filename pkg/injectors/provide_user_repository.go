package injectors

import (
	"github.com/oncelyunus/go_boilerplate/pkg/app/repository/user"
	"github.com/oncelyunus/go_boilerplate/pkg/internal"
)

func ProvideUserRepository(baseMongoRepo *internal.BaseMongoRepo) user.UserRepository {
	return user.NewUserRepository(baseMongoRepo)
}
