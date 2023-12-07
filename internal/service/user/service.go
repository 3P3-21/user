package user

import (
	"github.com/3P3-21/user/internal/server/req"
	"github.com/3P3-21/user/internal/store"
)

type User interface {
	SignUp(ctx *req.Ctx) error
	SignIn(ctx *req.Ctx) error
}

type Service struct {
	userStore store.UserStore
}

func New(userStore store.UserStore) *Service {
	return &Service{
		userStore: userStore,
	}
}
