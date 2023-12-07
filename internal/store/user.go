package store

import (
	"time"
)

type User struct {
	ID        int
	Email     string
	Password  string
	FirstName string
	LastName  string
	RoleID    uint16
	CreatedAt time.Time
}

type SignUpOpts struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
}

type GetByEmailOpts struct {
	Email string
}

type UserStore interface {
	SaveUser(opts SignUpOpts) (User, error)
}
