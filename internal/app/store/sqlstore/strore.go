package sqlstore

import (
	"ApiServer/internal/app/store"
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	db                *sql.DB
	studentRepository *StudentRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Student() store.StudentRepository {
	if s.studentRepository != nil {
		return s.studentRepository
	}

	s.studentRepository = &StudentRepository{
		store: s,
	}

	return s.studentRepository
}
