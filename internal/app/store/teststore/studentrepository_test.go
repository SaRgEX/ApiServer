package teststore_test

import (
	"ApiServer/internal/app/model"
	"ApiServer/internal/app/store"
	"ApiServer/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStudentRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestStudent(t)
	assert.NoError(t, s.Student().Create(u))
	assert.NotNil(t, u)
}

func TestStudentRepository_FindByLogin(t *testing.T) {
	s := teststore.New()
	stud := model.TestStudent(t)
	_, err := s.Student().FindByLogin(stud.Login)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.Student().Create(stud)

	stud, err = s.Student().FindByLogin(stud.Login)
	assert.NoError(t, err)
	assert.NotNil(t, stud)
}

func TestStudentRepository_Find(t *testing.T) {
	s := teststore.New()
	stud := model.TestStudent(t)

	s.Student().Create(stud)

	stud, err := s.Student().Find(stud.ID)
	assert.NoError(t, err)
	assert.NotNil(t, stud)
}
