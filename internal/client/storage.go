package client

import (
	"database/sql"
	"errors"
)

type ClientStorage interface {
	FindByUserID(idUser int) ([]Client, error)
	FindByID(id int) (*Client, error)
	Create(request *CreateClientRequest, idUser int) (int, error)
	Delete(id int) error
	Update(id int, request *CreateClientRequest) error
}

type PostgresClientStorage struct {
	db *sql.DB
}

func NewPostgresClientStorage(db *sql.DB) *PostgresClientStorage {
	return &PostgresClientStorage{
		db: db,
	}
}

func (s *PostgresClientStorage) FindByUserID(idUser int) ([]Client, error) {
	clients := []Client{}
	client := Client{}

	rows, err := s.db.Query("SELECT id, name, whatsapp, email, created_at, updated_at FROM clients WHERE user_id = $1", idUser)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&client.ID, &client.Name, &client.Whatsapp, &client.Email, &client.CreatedAt, &client.UpdatedAt)
		if err != nil {
			return nil, err
		}

		clients = append(clients, client)
	}

	return clients, nil
}

func (s *PostgresClientStorage) FindByID(id int) (*Client, error) {
	client := Client{}

	rows, err := s.db.Query("SELECT id, name, whatsapp, email, created_at, updated_at FROM clients WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&client.ID, &client.Name, &client.Whatsapp, &client.Email, &client.CreatedAt, &client.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &client, nil
}

func (s *PostgresClientStorage) Create(request *CreateClientRequest, idUser int) (int, error) {
	id := 0

	row := s.db.QueryRow(
		`INSERT INTO clients (name, whatsapp, email, user_id) VALUES ($1, $2, $3, $4) RETURNING id`,
		request.Name, request.Whatsapp, request.Email, idUser,
	).Scan(&id)

	if row != nil {
		return 0, errors.New(row.Error())
	}

	return id, nil
}

func (s *PostgresClientStorage) Delete(id int) error {
	rows, err := s.db.Query("DELETE FROM clients WHERE id = $1", id)
	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}

func (s *PostgresClientStorage) Update(id int, request *CreateClientRequest) error {
	row := s.db.QueryRow(
		`UPDATE clients SET name = $1, whatsapp = $2, email = $3 WHERE id = $4`,
		request.Name, request.Whatsapp, request.Email, id,
	)

	if row.Err() != nil {
		return errors.New(row.Err().Error())
	}

	return nil
}
