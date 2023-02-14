package charge

import (
	"database/sql"
	"errors"
)

type ChargeStorage interface {
	FindByUserID(idUser int) ([]Charge, error)
	FindByID(id int) (*Charge, error)
	Create(request *CreateChargeRequest, idUser int) (int, error)
	Delete(id int) error
	Update(id int, request *CreateChargeRequest) error
}

type PostgresChargeStorage struct {
	db *sql.DB
}

func NewPostgresChargeStorage(db *sql.DB) *PostgresChargeStorage {
	return &PostgresChargeStorage{
		db: db,
	}
}

func (s *PostgresChargeStorage) FindByUserID(idUser int) ([]Charge, error) {
	charges := []Charge{}
	charge := Charge{}

	rows, err := s.db.Query("SELECT id, value, observation, notification_date, client_id, created_at, updated_at FROM charges WHERE user_id = $1", idUser)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&charge.ID, &charge.Value, &charge.Observation, &charge.NotificationDate, &charge.ClientId, &charge.CreatedAt, &charge.UpdatedAt)
		if err != nil {
			return nil, err
		}

		charges = append(charges, charge)
	}

	return charges, nil
}

func (s *PostgresChargeStorage) FindByID(id int) (*Charge, error) {
	charge := Charge{}

	rows, err := s.db.Query("SELECT id, value, observation, notification_date, client_id, created_at, updated_at FROM charges WHERE user_id = $1", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&charge.ID, &charge.Value, &charge.Observation, &charge.NotificationDate, &charge.ClientId, &charge.CreatedAt, &charge.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &charge, nil
}

func (s *PostgresChargeStorage) Create(request *CreateChargeRequest, idUser int) (int, error) {
	id := 0

	row := s.db.QueryRow(
		`INSERT INTO charges (value, observation, notification_date, client_id, user_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		request.Value, request.Observation, request.NotificationDate, request.ClientId, idUser,
	).Scan(&id)

	if row != nil {
		return 0, errors.New(row.Error())
	}

	return id, nil
}

func (s *PostgresChargeStorage) Delete(id int) error {
	rows, err := s.db.Query("DELETE FROM charges WHERE id = $1", id)
	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}

func (s *PostgresChargeStorage) Update(id int, request *CreateChargeRequest) error {
	row := s.db.QueryRow(
		`UPDATE charges SET value = $1, observation = $2, notification_date = $3, client_id = $4 WHERE id = $5`,
		request.Value, request.Observation, request.NotificationDate, request.ClientId, id,
	)

	if row.Err() != nil {
		return errors.New(row.Err().Error())
	}

	return nil
}
