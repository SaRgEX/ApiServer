package teststore

import (
	"ApiServer/internal/app/model"
	"ApiServer/internal/app/store"
)

type Store struct {
	studentRepository *StudentRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) Student() store.StudentRepository {
	if s.studentRepository != nil {
		return s.studentRepository
	}

	s.studentRepository = &StudentRepository{
		store:   s,
		student: make(map[int]*model.Student),
	}

	return s.studentRepository
}
