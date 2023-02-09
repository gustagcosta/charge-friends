package user

import (
	"context"
	"database/sql"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{
		db: db,
	}
}

func (s *UserStorage) CreateNewUser(newUser UserCreateRequest, ctx context.Context) error {
	stmt, err := s.db.Prepare("INSERT INTO users (name, corporate_name, email, password) VALUES ($1,$2,$3,$4)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(newUser.Name, newUser.CorporateName, newUser.Email, newUser.Password)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserStorage) FindUserByEmail(email string) (*User, error) {
	u := User{}

	rows, err := s.db.Query("SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if !rows.Next() {
		return &u, nil
	}

	err = rows.Scan(&u.ID, &u.Name, &u.CorporateName, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
