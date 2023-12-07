package user

import (
	"github.com/3P3-21/user/internal/server/req"
	"github.com/3P3-21/user/internal/store"
	"github.com/3P3-21/user/pkg/errs"
	"golang.org/x/crypto/bcrypt"
)

type SignUpRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6,max=32"`
	FirstName string `json:"firstName" validate:"required,min=6,max=32"`
	LastName  string `json:"lastName" validate:"required,min=6,max=32"`
}

type SignUnResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (s *Service) SignUp(ctx *req.Ctx) error {
	var request SignUpRequest

	err := ctx.ParseJSON(&request)
	if err != nil {
		errs.ReturnError(ctx.Writer, errs.InvalidJSON)
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		errs.ReturnError(ctx.Writer, errs.InternalServerError)
	}

	// When the database query part is ready, store.User will be returned here to generate the JWT
	_, err = s.userStore.SaveUser(store.SignUpOpts{
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Password:  string(passwordHash),
	})
	if err != nil {
		errs.ReturnError(ctx.Writer, errs.InternalServerError)
	}

	// TODO: impl the JWT generation system
	return ctx.JSON(SignUnResponse{})
}
