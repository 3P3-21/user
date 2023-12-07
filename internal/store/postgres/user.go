package postgres

import (
	"database/sql"

	"github.com/3P3-21/curriculum/internal/store"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) store.UserStore {
	return &UserStore{
		db: db,
	}
}

func (s *UserStore) SaveUser(opts store.SignUpOpts) (store.User, error) {
	var user store.User

	// TODO: Write SQL func
	resp := s.db.QueryRow("SELECT * FROM create_user($1, $2, $3)",
		opts.Email, opts.FirstName, opts.LastName, opts.Password)

	err := resp.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Password, &user.RoleID, &user.CreatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *UserStore) GetUserByEmail(opts store.GetByEmailOpts) (store.User, error) {
	var user store.User

	// TODO: Write SQL func
	resp := s.db.QueryRow("SELECT * FROM get_user_by_email($1)", opts.Email)

	err := resp.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Password, &user.RoleID, &user.CreatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}
