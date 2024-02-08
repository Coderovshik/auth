package domain

import "context"

type RegisterUsecase interface {
	Register(ctx context.Context, info UserInfo) (string, error)
}
