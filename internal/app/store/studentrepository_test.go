package store_test

import (
	"ApiServer/internal/app/model"
	"ApiServer/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStudentRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("student")

	u, err := s.Student().Create(model.TestStudent(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestStudentRepository_FindByLogin(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("student")
	login := "login"
	_, err := s.Student().FindByLogin(login)
	assert.Error(t, err)

	stud := model.TestStudent(t)
	stud.Login = login

	s.Student().Create(stud)
	stud, err = s.Student().FindByLogin(login)
	assert.NoError(t, err)
	assert.NotNil(t, stud)
}
