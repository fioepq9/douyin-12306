package repository

import (
	"context"
	"douyin-12306/models"
	"douyin-12306/pkg/util"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestUserDAO_Register(t *testing.T) {
	var user models.User

	var (
		username = "username101"
		password = "password101"
		name     = "username101"
	)

	ctx := context.Background()
	userPointer, err := NewUserDAOInstance().Register(ctx, username, password, name)
	t.Log(userPointer, err)
	id, err := R.Redis.Get(ctx, models.User{}.UsernameKeyPrefix()+username).Int64()
	t.Log(id, err)
	err = R.Redis.Get(ctx, models.User{}.KeyPrefix()+util.Int64ToString(id)).Scan(&user)
	t.Log(user, err)
}
