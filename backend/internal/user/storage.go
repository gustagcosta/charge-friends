package user

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type UserStorage struct {
	db *sql.DB
}
type User struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	CorporateName string    `json:"corporate_name"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type UserCreateRequest struct {
	Name          string `json:"name"`
	CorporateName string `json:"corporate_name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserCreateRequest) Validate() error {
	if u.Name == "" {
		return errors.New("a user must have a name")
	}

	if u.CorporateName == "" {
		return errors.New("a user must have a corporate name")
	}

	if u.Email == "" {
		return errors.New("a user must have a email")
	}

	if u.Password == "" {
		return errors.New("a user must have a password")
	}

	if len(u.Password) < 6 {
		return errors.New("password must have at least 6 characteres")
	}

	return nil
}

func (u *UserLoginRequest) Validate() error {
	if u.Email == "" {
		return errors.New("email not provided")
	}

	if u.Password == "" {
		return errors.New("password not provided")
	}

	return nil
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
