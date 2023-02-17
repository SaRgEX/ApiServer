package sqlstore_test

import (
	"ApiServer/internal/app/model"
	"ApiServer/internal/app/store"
	"ApiServer/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStudentRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("student")

	s := sqlstore.New(db)
	student := model.TestStudent(t)
	assert.NoError(t, s.Student().Create(student))
	assert.NotNil(t, student)
}

func TestStudentRepository_FindByLogin(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("student")
	stud := model.TestStudent(t)

	s := sqlstore.New(db)
	_, err := s.Student().FindByLogin(stud.Login)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.Student().Create(stud)

	stud, err = s.Student().FindByLogin(stud.Login)
	assert.NoError(t, err)
	assert.NotNil(t, stud)
}

func TestStudentRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("student")
	stud := model.TestStudent(t)

	s := sqlstore.New(db)

	s.Student().Create(stud)

	stud, err := s.Student().Find(stud.ID)
	assert.NoError(t, err)
	assert.NotNil(t, stud)
}
