package domain

import "context"

type UserInfo struct {
	Email    string
	Password string
}

type User struct {
	ID       string
	Email    string
	PassHash []byte
}

type UserRepository interface {
	Save(ctx context.Context, email string, passHash []byte) (string, error)
	Get(ctx context.Context, email string) (User, error)
	IsAdmin(ctx context.Context, id string) (bool, error)
}
