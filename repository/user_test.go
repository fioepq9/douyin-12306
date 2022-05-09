package repository

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestUserDAO_Register(t *testing.T) {
	id, err := NewUserDAOInstance().Register("username", "password", "username")
	t.Log(id, err)
}
