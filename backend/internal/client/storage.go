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

func (s *ClientStorage) FindClients(id int) ([]Client, error) {
	clients := []Client{}
	client := Client{}

	rows, err := s.db.Query("SELECT id, name, whatsapp, email, created_at, updated_at FROM clients WHERE user_id = $1", id)
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

func (s *ClientStorage) CreateNewClient(newClient *ClientCreateRequest, idUser int) (int, error) {
	id := 0

	row := s.db.QueryRow(
		`INSERT INTO clients (name, whatsapp, email, user_id) VALUES ($1, $2, $3, $4) RETURNING id`,
		newClient.Name, newClient.Whatsapp, newClient.Email, idUser,
	).Scan(&id)

	if row != nil {
		return 0, errors.New(row.Error())
	}

	return id, nil
}
