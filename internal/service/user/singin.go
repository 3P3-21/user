package user

import (
	"github.com/3P3-21/user/internal/server/req"
	"github.com/3P3-21/user/internal/store"
	"github.com/3P3-21/user/pkg/errs"
	"golang.org/x/crypto/bcrypt"
)

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=32"`
}

type SignInResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (s *Service) SignIn(ctx *req.Ctx) error {
	var request SignUpRequest

	err := ctx.ParseJSON(&request)
	if err != nil {
		errs.ReturnError(ctx.Writer, errs.InvalidJSON)
	}

	user, err := s.userStore.SaveUser(store.SignUpOpts{
		Email: request.Email,
	})

	err = bcrypt.CompareHashAndPassword([]byte(request.Password), []byte(user.Password))
	if err != nil {
		errs.ReturnError(ctx.Writer, errs.IncorrectLoginOrPassword)
	}

	// TODO: impl the JWT generation system
	return ctx.JSON(SignUnResponse{})
}
