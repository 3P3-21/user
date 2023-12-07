package service

import (
	"github.com/3P3-21/curriculum/internal/service/user"
	"github.com/3P3-21/curriculum/internal/store"
)

type Service struct {
	User user.User
}

func New(userStore store.UserStore) *Service {
	return &Service{
		User: user.New(userStore),
	}
}
