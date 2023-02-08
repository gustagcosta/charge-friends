package user

import (
	"context"
	"database/sql"
	"time"
)

// how the user is stored in the database
type userDb struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	CorporateName string    `json:"corporate_name"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{
		db: db,
	}
}

func (s *UserStorage) getAllUsers(ctx context.Context) ([]userDb, error) {
	query, err := s.db.Query("SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	users := []userDb{}

	for query.Next() {
		u := userDb{}
		err := query.Scan(&u.ID, &u.Name, &u.CorporateName, &u.Password, &u.Email, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}
