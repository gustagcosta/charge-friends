package client

import (
	"database/sql"
	"errors"
)

type ClientStorage struct {
	db *sql.DB
}

func NewClientStorage(db *sql.DB) *ClientStorage {
	return &ClientStorage{
		db: db,
	}
}

func (s *ClientStorage) FindClients(idUser int) ([]Client, error) {
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

func (s *ClientStorage) FindClientByID(idClient int) (*Client, error) {
	client := Client{}

	rows, err := s.db.Query("SELECT id, name, whatsapp, email, created_at, updated_at FROM clients WHERE id = $1", idClient)
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

func (s *ClientStorage) CreateNewClient(clientCreateRequest *ClientCreateRequest, idUser int) (int, error) {
	id := 0

	row := s.db.QueryRow(
		`INSERT INTO clients (name, whatsapp, email, user_id) VALUES ($1, $2, $3, $4) RETURNING id`,
		clientCreateRequest.Name, clientCreateRequest.Whatsapp, clientCreateRequest.Email, idUser,
	).Scan(&id)

	if row != nil {
		return 0, errors.New(row.Error())
	}

	return id, nil
}

func (s *ClientStorage) DeleteClientByID(idClient int) error {
	rows, err := s.db.Query("DELETE FROM clients WHERE id = $1", idClient)
	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}

func (s *ClientStorage) UpdateClient(idClient int, clientCreateRequest *ClientCreateRequest) error {
	row := s.db.QueryRow(
		`UPDATE clients SET name = $1, whatsapp = $2, email = $3 WHERE id = $4`,
		clientCreateRequest.Name, clientCreateRequest.Whatsapp, clientCreateRequest.Email, idClient,
	)

	if row.Err() != nil {
		return errors.New(row.Err().Error())
	}

	return nil
}
