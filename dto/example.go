package dto

import "github.com/3P3-21/curriculum/model"

type User struct {
	ID        int
	Email     string
	FirstName string
	LastName  string
	RoleID    uint16
}

func NewUser(m model.User) User {
	return User{
		ID:        m.ID,
		Email:     m.Email,
		FirstName: m.FirstName,
		LastName:  m.LastName,
		RoleID:    m.RoleID,
	}
}

func NewExamples(m []model.User) []User {
	var users []User
	for i := range m {
		users = append(users, NewUser(m[i]))
	}
	return users
}
