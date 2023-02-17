package teststore

import (
	"ApiServer/internal/app/model"
	"ApiServer/internal/app/store"
)

type StudentRepository struct {
	store   *Store
	student map[int]*model.Student
}

func (r *StudentRepository) Create(s *model.Student) error {
	if err := s.Validate(); err != nil {
		return err
	}

	if err := s.BeforeCreate(); err != nil {
		return err
	}

	s.ID = len(r.student) + 1
	r.student[s.ID] = s

	return nil
}

func (r *StudentRepository) FindByLogin(login string) (*model.Student, error) {
	for _, s := range r.student {
		if s.Login == login {
			return s, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

func (r *StudentRepository) Find(id int) (*model.Student, error) {
	s, ok := r.student[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return s, nil
}
