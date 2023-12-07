package model

import (
	"time"

	"github.com/3P3-21/curriculum/internal/store"
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

func NewUser(s store.User) User {
	user := User{
		ID:        s.ID,
		Email:     s.Email,
		Password:  s.Password,
		FirstName: s.FirstName,
		LastName:  s.LastName,
		RoleID:    s.RoleID,
		CreatedAt: s.CreatedAt,
	}

	return user
}

func NewUsers(s []store.User) []User {
	var users []User
	for i := range s {
		users = append(users, NewUser(s[i]))
	}
	return users
}
