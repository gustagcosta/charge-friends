package user

import (
	"errors"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	PixKey    string    `json:"pix_key"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserCreateRequest struct {
	Name     string `json:"name"`
	PixKey   string `json:"pix_key"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserCreateRequest) Validate() error {
	if u.Name == "" {
		return errors.New("a user must have a name")
	}

	if u.PixKey == "" {
		return errors.New("a user must have a pix key")
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

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
