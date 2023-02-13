package store

import "ApiServer/internal/app/model"

type StudentRepository struct {
	store *Store
}

func (r *StudentRepository) Create(s *model.Student) (*model.Student, error) {
	if err := s.Validate(); err != nil {
		return nil, err
	}

	if err := s.BeforeCreate(); err != nil {
		return nil, err
	}

	if err := r.store.db.QueryRow(
		"INSERT INTO student (login, encrypted_password) VALUES ($1, $2) RETURNING id",
		s.Login,
		s.EncryptedPassword,
	).Scan(&s.ID); err != nil {
		return nil, err
	}

	return s, nil
}

func (r *StudentRepository) FindByLogin(login string) (*model.Student, error) {
	s := &model.Student{}
	if err := r.store.db.QueryRow(
		"SELECT id, login, encrypted_password FROM student WHERE login = $1",
		login,
	).Scan(
		&s.ID,
		&s.Login,
		&s.EncryptedPassword,
	); err != nil {
		return nil, err
	}
	return s, nil
}
