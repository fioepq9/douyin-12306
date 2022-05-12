package repository

import (
	"context"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestUserDAO_Register(t *testing.T) {
	ctx := context.Background()
	id, err := NewUserDAOInstance().Register(ctx, "username101", "password101", "username101")
	t.Log(id, err)
}
