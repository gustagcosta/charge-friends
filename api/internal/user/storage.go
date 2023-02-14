package user

import (
	"database/sql"
)

type UserStorage interface {
	FindByEmail(email string) (*User, error)
	Create(request *UserCreateRequest) error
}

type PostgresUserStorage struct {
	db *sql.DB
}

func NewPostgresUserStorage(db *sql.DB) *PostgresUserStorage {
	return &PostgresUserStorage{
		db: db,
	}
}

func (s *PostgresUserStorage) Create(request *UserCreateRequest) error {
	stmt, err := s.db.Prepare("INSERT INTO users (name, pix_key, email, password) VALUES ($1,$2,$3,$4)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(request.Name, request.PixKey, request.Email, request.Password)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresUserStorage) FindByEmail(email string) (*User, error) {
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
