package model

import "testing"

func TestStudent(t *testing.T) *Student {
	return &Student{
		Login:    "valid",
		Password: "password",
	}
}
