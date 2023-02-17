package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

type Student struct {
	ID                int    `json:"id"`
	Login             string `json:"login"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
}

func (s *Student) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.Login, validation.Required, validation.Length(4, 24)),
		validation.Field(&s.Password, validation.By(requiredIf(s.EncryptedPassword == "")), validation.Length(8, 100)),
	)
}

func (s *Student) BeforeCreate() error {
	if len(s.Password) > 0 {
		enc, err := encryptString(s.Password)
		if err != nil {
			return err
		}
		s.EncryptedPassword = enc
	}
	return nil
}

func (s *Student) Sanitaize() {
	s.Password = ""
}

func (s *Student) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(s.EncryptedPassword), []byte(password)) == nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
