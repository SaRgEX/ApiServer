package store

import "ApiServer/internal/app/model"

type StudentRepository interface {
	Create(student *model.Student) error
	Find(int) (*model.Student, error)
	FindByLogin(string) (*model.Student, error)
}
