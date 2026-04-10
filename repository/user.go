package repository

import (
	"errors"

	"github.com/chiayu/auth-api/model"
)

type UserRepository struct {
	users  []model.User
	nextID uint
}

func NewUserRepository() *UserRepository {
	return &UserRepository{nextID: 1}
}

func (r *UserRepository) Save(user model.User) (model.User, error) {
	user.ID = r.nextID
	r.nextID++
	r.users = append(r.users, user)
	return user, nil
}

func (r *UserRepository) FindByUsername(username string) (model.User, error) {
	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}
	return model.User{}, errors.New("user not found")
}

func (r *UserRepository) ExistsByUsername(username string) bool {
	for _, user := range r.users {
		if user.Username == username {
			return true
		}
	}
	return false
}
