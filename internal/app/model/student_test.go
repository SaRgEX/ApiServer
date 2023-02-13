package model_test

import (
	"ApiServer/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStudent_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		s       func() *model.Student
		isValid bool
	}{
		{
			name: "valid",
			s: func() *model.Student {
				return model.TestStudent(t)
			},
			isValid: true,
		},
		{
			name: "empty login",
			s: func() *model.Student {
				s := model.TestStudent(t)
				s.Login = ""
				return s
			},
			isValid: false,
		},
		{
			name: "invalid login",
			s: func() *model.Student {
				s := model.TestStudent(t)
				s.Login = "inv"
				return s
			},
			isValid: false,
		},
		{
			name: "empty password",
			s: func() *model.Student {
				s := model.TestStudent(t)
				s.Password = ""
				return s
			},
			isValid: false,
		},
		{
			name: "short password",
			s: func() *model.Student {
				s := model.TestStudent(t)
				s.Password = "short"
				return s
			},
			isValid: false,
		},
		{
			name: "with encrypt password",
			s: func() *model.Student {
				s := model.TestStudent(t)
				s.Password = ""
				s.EncryptedPassword = "encrypted_password"
				return s
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.s().Validate())
			} else {
				assert.Error(t, tc.s().Validate())
			}
		})
	}
}

func TestStudent_BeforeCreate(t *testing.T) {
	s := model.TestStudent(t)
	assert.NoError(t, s.BeforeCreate())
	assert.NotNil(t, s.EncryptedPassword)
}
