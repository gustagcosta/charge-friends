package user

import (
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

func (s *UserStorage) CreateNewUser(newUser UserCreateRequest) error {
	stmt, err := s.db.Prepare("INSERT INTO users (name, pix_key, email, password) VALUES ($1,$2,$3,$4)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(newUser.Name, newUser.PixKey, newUser.Email, newUser.Password)
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

	err = rows.Scan(&u.ID, &u.Name, &u.PixKey, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (s *UserStorage) FindUserByID(id int) (*User, error) {
	u := User{}

	rows, err := s.db.Query("SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if !rows.Next() {
		return &u, nil
	}

	err = rows.Scan(&u.ID, &u.Name, &u.PixKey, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}