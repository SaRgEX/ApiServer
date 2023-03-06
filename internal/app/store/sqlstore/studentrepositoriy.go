package sqlstore

import (
	"ApiServer/internal/app/model"
	"ApiServer/internal/app/store"
	"database/sql"
)

type StudentRepository struct {
	store *Store
}

func (r *StudentRepository) Create(s *model.Student) error {
	if err := s.Validate(); err != nil {
		return err
	}

	if err := s.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO student (login, encrypted_password, name, surname, patronymic, address, group_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		s.Login,
		s.EncryptedPassword,
		s.Name,
		s.Surname,
		s.Patronymic,
		s.Address,
		s.GroupId,
	).Scan(&s.ID)
}

func (r *StudentRepository) FindByLogin(login string) (*model.Student, error) {
	s := &model.Student{}
	if err := r.store.db.QueryRow(
		"SELECT id, login, encrypted_password, group_id FROM student WHERE login = $1",
		login,
	).Scan(
		&s.ID,
		&s.Login,
		&s.EncryptedPassword,
		&s.GroupId,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return s, nil
}

func (r *StudentRepository) Find(id int) (*model.Student, error) {
	s := &model.Student{}
	if err := r.store.db.QueryRow(
		"SELECT id, login, encrypted_password, group_id FROM student WHERE id = $1",
		id,
	).Scan(
		&s.ID,
		&s.Login,
		&s.EncryptedPassword,
		&s.GroupId,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return s, nil
}

func (r *StudentRepository) FindGroup() (*model.Group, error) {
	group := &model.Group{}
	if err := r.store.db.QueryRow(
		"SELECT id FROM \"group\"",
	).Scan(
		&group.Id,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return group, nil
}
